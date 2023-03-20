package models

import (
	"github.com/google/uuid"
)

type ContextValues struct {
	Account       uuid.UUID
	SessionClaims *SessionClaims
}

type ContextKey string

var (
	ContextValuesKey = ContextKey("ContextValues")
)
