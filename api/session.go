package api

import (
	"net/http"

	"github.com/effectindex/tripreporter/types"
	"github.com/gorilla/mux"
)

// SessionPost path is /api/v1/session
func SessionPost(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)

	ctx.Handle(w, r, types.ErrorApiNotImplemented)
}

// SessionGet path is /api/v1/session/{id}
func SessionGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		ctx.Handle(w, r, types.ErrorApiSessionNilId)
		return
	}

	ctx.Logger.Infow("Got ID", "id", id)
}
