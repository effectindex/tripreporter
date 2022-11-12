package models

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Unique struct {
	ID uuid.UUID `json:"id"`
}

func (u *Unique) InitUUID(logger *zap.SugaredLogger) error {
	if u.NilUUID() {
		var err error
		u.ID, err = uuid.NewUUID()
		if err != nil {
			logger.Warnw("Failed to make UUID", zap.Error(err))
			return err
		}
	}

	return nil
}

func (u *Unique) NilUUID() bool {
	return &u.ID == nil || u.ID == uuid.Nil
}
