// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/effectindex/tripreporter/types"
	"github.com/effectindex/tripreporter/util"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Session struct { // TODO: Really, this should be keeping track of IP and useragent in the future.
	types.Context
	Unique
	Index   int    `json:"-" db:"session_index"`
	Key     Unique `json:"-" db:"session_key"`
	Refresh string `json:"-" db:"-"` // Stored in Redis only (s.Cache).
}

type SessionClaims struct {
	jwt.RegisteredClaims
	Account uuid.NullUUID `json:"account_id"`
	Session uuid.NullUUID `json:"session_id"`
}

// Get will return all sessions matching an account ID, and can be empty.
// This does not include the refresh token.
func (s *Session) Get() ([]*Session, error) {
	s.InitType(s)
	db := s.DB()
	defer db.Commit(context.Background())

	s1 := make([]*Session, 0)

	// You should specify the account ID here to go with it
	if s.NilUUID() {
		return s1, types.ErrorSessionNotSpecified
	}

	if err := pgxscan.Select(context.Background(), db, &s1,
		`select * from sessions where account_id = $1;`, s.ID,
	); err != nil {
		s.Logger.Warnw("Failed to get sessions from DB", zap.Error(err))
		return s1, err
	}

	return s1, nil
}

// GetByKey will return the first matching session by the set session key, and will return an error if not found.
// This includes the index, not the refresh token.
func (s *Session) GetByKey() (*Session, error) {
	s.InitType(s)

	if s.Key.NilUUID() { // you need to specify key, it's faster to check here
		return s, types.ErrorSessionKeyNotFound
	}

	sessions, err := s.Get()
	if err != nil {
		return s, err
	}

	for _, session := range sessions {
		if session.Key.ID == s.Key.ID {
			s.Index = session.Index
			return s, nil
		}
	}

	return s, types.ErrorSessionIndexNotFound
}

// Post creates a new session for an account. Returns the session key and refresh token, but not the index.
func (s *Session) Post() (*Session, error) {
	s.InitType(s)
	db := s.DB()
	defer db.Commit(context.Background())

	// You should specify the account ID here to go with it
	if s.NilUUID() {
		return s, types.ErrorSessionNotSpecified
	}

	if s.Key.NilUUID() { // UUID v4 is "fine" here, really.
		if err := s.Key.InitUUIDv4(s.Logger); err != nil {
			return s, err
		}
	}

	if _, err := db.Exec(context.Background(),
		`insert into sessions(account_id, session_index, session_key) (select $1, coalesce(max(sessions.session_index), 0) + 1, $2 from sessions where account_id = $1);`,
		s.ID, s.Key.ID,
	); err != nil {
		s.Logger.Warnw("Failed to session to DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return s, err
	}

	// Now that we have a session in the database, lets make a refresh token for it.
	b, err := util.GenerateRandomBytes(16)
	if err != nil {
		s.Logger.Warnw("Failed to create bytes for refresh token", zap.Error(err))
		_ = db.Rollback(context.Background())
		return s, err
	}

	// Encode the refresh token to a b64 string, and set it to expire from Redis a week from now, with the refresh as the key and the account ID as the value
	s.Refresh = base64.StdEncoding.EncodeToString(b)
	err = s.Cache.Set(context.Background(), s.Refresh, s.ID.String(), time.Hour*24*7).Err()
	if err != nil {
		s.Logger.Warnw("Failed to set cache for refresh token", zap.Error(err))
		_ = db.Rollback(context.Background())
		return s, err
	}

	return s, nil
}

// Delete will invalidate all sessions for an account
func (s *Session) Delete() (*Session, error) {
	s.InitType(s)
	db := s.DB()
	defer db.Commit(context.Background())

	// You should specify the account ID here to go with it
	if s.NilUUID() {
		return s, types.ErrorSessionNotSpecified
	}

	// Now delete all sessions from the account
	if _, err := db.Exec(context.Background(), `delete from sessions where account_id=$1;`, s.ID); err != nil {
		s.Logger.Warnw("Failed to delete account sessions from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return s, err
	}

	// TODO: Part of refactoring to pointer-based model
	return s.ClearAll(), nil
}

// DeleteByKey will delete just one session from an account
func (s *Session) DeleteByKey() (*Session, error) {
	s.InitType(s)
	db := s.DB()
	defer db.Commit(context.Background())

	// You should specify the account ID here to go with it
	if s.NilUUID() {
		return s, types.ErrorSessionNotSpecified
	}

	if s.Key.NilUUID() { // you need to specify key, it's faster to check here
		return s, types.ErrorSessionKeyNotFound
	}

	// Now delete the matching session from the account
	if _, err := db.Exec(context.Background(), `delete from sessions where account_id=$1 and session_key=$2;`, s.ID, s.Key.ID); err != nil {
		s.Logger.Warnw("Failed to delete individual session from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return s, err
	}

	// TODO: Part of refactoring to pointer-based model
	return s.ClearAll(), nil
}

func (s *Session) ClearAll() *Session {
	s.InitType(s)
	return &Session{Context: s.Context, Unique: s.Unique}
}
