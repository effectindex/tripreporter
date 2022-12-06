package api

import (
	"github.com/effectindex/tripreporter/types"
	"net/http"
)

type Message int64

const (
	MsgUnknown Message = iota
	MsgNotImplemented
	MsgMethodNotAllowed
	MsgInvalidEndpoint
	MsgInvalidApiVersion
	MsgSessionNilId
)

func (e Message) String() string {
	switch e {
	case MsgNotImplemented:
		return types.ErrorNotImplemented.Error()
	case MsgMethodNotAllowed:
		return "Method not allowed!"
	case MsgInvalidEndpoint:
		return "Invalid endpoint!"
	case MsgInvalidApiVersion:
		return "Invalid API version!"
	case MsgSessionNilId:
		return "`id` is nil or unset!"
	default:
		return types.ErrorGenericUnknown.Error()
	}
}

func (e Message) Status() int {
	switch e {
	case MsgNotImplemented:
		return http.StatusNotImplemented
	case MsgMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case MsgInvalidEndpoint:
		return http.StatusBadRequest
	case MsgInvalidApiVersion:
		return http.StatusBadGateway
	case MsgSessionNilId:
		return http.StatusBadRequest
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
