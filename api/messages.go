package api

import (
	"net/http"

	"github.com/effectindex/tripreporter/types"
)

type Message int64

const (
	MsgUnknown Message = iota
	MsgNotImplemented
	MsgNotFound
	MsgMethodNotAllowed
	MsgInvalidEndpoint
	MsgInvalidApiVersion
	MsgNilVariable
	MsgOk
)

func (e Message) String() string {
	switch e {
	case MsgNotImplemented:
		return types.ErrorNotImplemented.Error()
	case MsgNotFound:
		return http.StatusText(http.StatusNotFound)
	case MsgMethodNotAllowed:
		return "Method not allowed!"
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
	case MsgNotFound:
		return http.StatusNotFound
	case MsgMethodNotAllowed:
		return http.StatusMethodNotAllowed
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
