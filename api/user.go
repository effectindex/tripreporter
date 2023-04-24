// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package api

import (
	"net/http"

	"github.com/effectindex/tripreporter/models"
	"github.com/google/uuid"
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
	vars := mux.Vars(r)
	idStr, _ := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.Handle(w, r, MsgNotFound)
		return
	}

	user, err := (&models.User{Context: ctx.Context, Unique: models.Unique{ID: id}}).Get()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.HandleJson(w, r, user.CopyPublic(), http.StatusOK)
}

// UserPatch path is /api/v1/user/{id}
func UserPatch(w http.ResponseWriter, r *http.Request) {
	ctx.Handle(w, r, MsgNotImplemented)
}
