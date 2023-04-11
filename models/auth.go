// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

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
