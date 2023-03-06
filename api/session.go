package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupSessionEndpoints(v1 *mux.Router) {
	a1 := v1.Methods(http.MethodGet, http.MethodPatch, http.MethodDelete).Subrouter()
	a1.Use(AuthMiddleware())

	v1.HandleFunc("/session", SessionPost).Methods(http.MethodPost)
	v1.HandleFunc("/session/{id}", SessionGet).Methods(http.MethodGet)
	a1.HandleFunc("/session/validate", SessionGetValidate).Methods(http.MethodGet)
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

// SessionGetValidate path is /api/v1/session/validate
func SessionGetValidate(w http.ResponseWriter, r *http.Request) {
	// This might look like it does nothing, but it is called with the AuthMiddleware, which means if you don't have
	// a valid session it will return a 403 from there, instead.
	ctx.Handle(w, r, MsgOk)
}
