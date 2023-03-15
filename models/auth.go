package models

type ContextValues struct {
	SessionClaims *SessionClaims
}

type ContextKey string

var (
	ContextValuesKey = ContextKey("ContextValues")
)
