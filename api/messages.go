// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package api

import (
	"net/http"

	"github.com/effectindex/tripreporter/types"
)

type Message int64

const (
	MsgUnknown Message = iota
	MsgNotImplemented
	MsgForbidden
	MsgNotFound
	MsgMethodNotAllowed
	MsgNotAcceptable
	MsgInvalidEndpoint
	MsgInvalidApiVersion
	MsgNilVariable
	MsgOk
)

func (e Message) String() string {
	switch e {
	case MsgNotImplemented:
		return types.ErrorNotImplemented.Error()
	case MsgForbidden:
		return http.StatusText(http.StatusForbidden)
	case MsgNotFound:
		return http.StatusText(http.StatusNotFound)
	case MsgMethodNotAllowed:
		return http.StatusText(http.StatusMethodNotAllowed)
	case MsgNotAcceptable:
		return http.StatusText(http.StatusNotAcceptable)
	case MsgInvalidEndpoint:
		return "Invalid API endpoint!"
	case MsgInvalidApiVersion:
		return "Invalid API version!"
	case MsgNilVariable:
		return " is nil or unset!"
	case MsgOk:
		return "Ok"
	default:
		return types.ErrorUnknown.Error()
	}
}

func (e Message) Status() int {
	switch e {
	case MsgNotImplemented:
		return http.StatusNotImplemented
	case MsgForbidden:
		return http.StatusForbidden
	case MsgNotFound:
		return http.StatusNotFound
	case MsgMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case MsgNotAcceptable:
		return http.StatusNotAcceptable
	case MsgInvalidEndpoint:
		return http.StatusBadRequest
	case MsgInvalidApiVersion:
		return http.StatusBadRequest
	case MsgNilVariable:
		return http.StatusBadRequest
	case MsgOk:
		return http.StatusOK
	default:
		return http.StatusInternalServerError
	}
}

func (e Message) Message() (string, int) {
	return e.String(), e.Status()
}

func (e Message) Error() string {
	return e.String()
}
