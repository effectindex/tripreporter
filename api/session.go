package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupSessionEndpoints(v1 *mux.Router) {
	v1.HandleFunc("/session", SessionPost).Methods(http.MethodPost)
	v1.HandleFunc("/session/{id}", SessionGet).Methods(http.MethodGet)
}

// SessionPost path is /api/v1/session
func SessionPost(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)

	ctx.Handle(w, r, MsgNotImplemented)
}

// SessionGet path is /api/v1/session/{id}
func SessionGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		ctx.Handle(w, r, MsgNilVariable)
		return
	}

	ctx.Logger.Infow("Got ID", "id", id)
}
