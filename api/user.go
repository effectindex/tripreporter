package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupUserEndpoints(v1 *mux.Router) {
	v1.HandleFunc("/user", UserPost).Methods(http.MethodPost)
	v1.HandleFunc("/user/{id}", UserGet).Methods(http.MethodGet)
	v1.HandleFunc("/user/{id}", UserPatch).Methods(http.MethodPatch)
}

// UserPost path is /api/v1/user
func UserPost(w http.ResponseWriter, r *http.Request) {
	ctx.Handle(w, r, MsgNotImplemented)
}

// UserGet path is /api/v1/user/{id}
func UserGet(w http.ResponseWriter, r *http.Request) {
	ctx.Handle(w, r, MsgNotImplemented)
}

// UserPatch path is /api/v1/user/{id}
func UserPatch(w http.ResponseWriter, r *http.Request) {
	ctx.Handle(w, r, MsgNotImplemented)
}
