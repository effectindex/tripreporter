// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package api

import (
	"net/http"

	"github.com/effectindex/tripreporter/models"
	"github.com/effectindex/tripreporter/types"
	"github.com/gorilla/mux"
)

func SetupSessionEndpoints(v1 *mux.Router) {
	a1 := v1.Methods(http.MethodGet, http.MethodPatch, http.MethodDelete).Subrouter()
	a1.Use(AuthMiddleware())

	v1.HandleFunc("/session", SessionPost).Methods(http.MethodPost)
	a1.HandleFunc("/session", SessionDelete).Methods(http.MethodDelete)
	a1.HandleFunc("/session/validate", SessionGetValidate).Methods(http.MethodGet)
}

// SessionPost path is /api/v1/session
func SessionPost(w http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)

	ctx.Handle(w, r, MsgNotImplemented)
}

// SessionDelete path is /api/v1/session
func SessionDelete(w http.ResponseWriter, r *http.Request) {
	ctxVals, ok := ctx.GetCtxValOrHandle(w, r)
	if !ok {
		return
	}

	if _, err := (&models.Session{
		Context: ctx.Context,
		Unique:  models.Unique{ID: ctxVals.Account},
		Key:     models.Unique{ID: ctxVals.SessionClaims.Session.UUID},
		Refresh: ctxVals.RefreshToken,
	}).DeleteByKey(); err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusForbidden)
		return
	}

	DeleteAuthCookies(w, types.CookieSessionID, types.CookieRefreshToken, types.CookieJwtToken)
	ctx.Handle(w, r, MsgOk)
}

// SessionGetValidate path is /api/v1/session/validate
func SessionGetValidate(w http.ResponseWriter, r *http.Request) {
	ctxVals, ok := ctx.GetCtxValOrHandle(w, r)
	if !ok {
		return
	}

	// This might look like it does nothing, but it is called with the AuthMiddleware, which means if you don't have
	// a valid session it will return a 403 from there, instead.
	ctx.Handle(w, r, MsgOk)
}
