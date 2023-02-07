package api

import (
	"net/http"
	"time"

	"github.com/effectindex/tripreporter/models"
	"github.com/effectindex/tripreporter/types"
	"github.com/effectindex/tripreporter/util"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func SetupAccountEndpoints(v1 *mux.Router) {
	a1 := v1.Methods(http.MethodGet, http.MethodPatch, http.MethodDelete).Subrouter()
	a1.Use(AuthMiddleware())

	v1.HandleFunc("/account", AccountPost).Methods(http.MethodPost)
	a1.HandleFunc("/account", AccountGet).Methods(http.MethodGet)
	a1.HandleFunc("/account", AccountPatch).Methods(http.MethodPatch)
	a1.HandleFunc("/account", AccountDelete).Methods(http.MethodDelete)
	v1.HandleFunc("/account/validate/email/{email}", AccountValidateEmail).Methods(http.MethodPost)
	v1.HandleFunc("/account/validate/username/{username}", AccountValidateUsername).Methods(http.MethodPost)
	v1.HandleFunc("/account/validate/password/{password}", AccountValidatePassword).Methods(http.MethodPost)
}

// AccountPost path is /api/v1/account
func AccountPost(w http.ResponseWriter, r *http.Request) {
	account, err := (&models.Account{Context: ctx.Context}).FromBody(r)
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	account = account.ClearImmutable()
	account.ID = uuid.Nil // We don't want to let users set the ID and so on when creating an account
	account, err = account.Post()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := (&models.Session{Context: ctx.Context, Unique: account.Unique}).Post()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	SetAuthCookie(w, util.CookieRefreshToken, session.Refresh, time.Now().Add(time.Minute*15)) // TODO: Change this once we've implemented refreshing

	ctx.HandleJson(w, r, account.ClearSensitive(), http.StatusCreated)
}

// AccountGet path is /api/v1/account
func AccountGet(w http.ResponseWriter, r *http.Request) {
	account, err := (&models.Account{Context: ctx.Context}).FromBody(r)
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	account, err = account.Get()
	if err != nil {
		if err == types.ErrorAccountNotSpecified || err == types.ErrorAccountNotFound {
			ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		} else {
			ctx.HandleStatus(w, r, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	ctx.HandleJson(w, r, account.ClearSensitive(), http.StatusOK)
}

// AccountPatch path is /api/v1/account
func AccountPatch(w http.ResponseWriter, r *http.Request) {
	account, err := (&models.Account{Context: ctx.Context}).FromBody(r)
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	account = account.ClearImmutable()
	account, err = account.Patch()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.HandleJson(w, r, account.ClearSensitive(), http.StatusOK)
}

// AccountDelete path is /api/v1/account
func AccountDelete(w http.ResponseWriter, r *http.Request) {
	account, err := (&models.Account{Context: ctx.Context}).FromBody(r)
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	account, err = account.Delete()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	DeleteAuthCookies(w, util.CookieRefreshToken, util.CookieJwtToken)

	ctx.HandleJson(w, r, account.ClearSensitive(), http.StatusOK)
}

// AccountValidateEmail path is /api/v1/account/validate/email/{email}
func AccountValidateEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, ok := vars["email"]

	if !ok {
		ctx.HandlePrefixed(w, r, "`email`", MsgNilVariable)
		return
	}

	_, err := (&models.Account{Context: ctx.Context, Email: email}).ValidateEmail()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.Handle(w, r, MsgOk)
}

// AccountValidateUsername path is /api/v1/account/validate/username/{username}
func AccountValidateUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]

	if !ok {
		ctx.HandlePrefixed(w, r, "`username`", MsgNilVariable)
		return
	}

	_, err := (&models.Account{Context: ctx.Context, Username: username}).ValidateUsername()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.Handle(w, r, MsgOk)
}

// AccountValidatePassword path is /api/v1/account/validate/password/{password}
func AccountValidatePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	password, ok := vars["password"]

	if !ok {
		ctx.HandlePrefixed(w, r, "`password`", MsgNilVariable)
		return
	}

	_, err := (&models.Account{Context: ctx.Context, Password: password}).ValidatePassword()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.Handle(w, r, MsgOk)
}
