package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Unique struct {
	ID   uuid.UUID `json:"id"`
	Type string    `json:"type"`
}

func (u *Unique) InitUUID(logger *zap.SugaredLogger) error {
	if u.NilUUID() {
		var err error
		u.ID, err = uuid.NewRandom()
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

func (u *Unique) InitType(a any) {
	if len(u.Type) == 0 {
		t := strings.Split(fmt.Sprintf("%T", a), ".")

		if len(t) > 0 {
			u.Type = strings.ToLower(t[len(t)-1])
		}
	}
}
