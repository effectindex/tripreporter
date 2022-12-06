package api

import (
	"github.com/effectindex/tripreporter/models"
	"github.com/effectindex/tripreporter/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupAccountEndpoints(v1 *mux.Router) {
	v1.HandleFunc("/account", AccountPost).Methods(http.MethodPost)
	v1.HandleFunc("/account/{id}", AccountGet).Methods(http.MethodGet)
}

// AccountPost path is /api/v1/account
func AccountPost(w http.ResponseWriter, r *http.Request) {
	//account, err := (&models.Account{}).Post()
	ctx.Handle(w, r, MsgNotImplemented)
}

// AccountGet path is /api/v1/account/{id}
func AccountGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		ctx.Handle(w, r, MsgSessionNilId)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := (&models.Account{Context: ctx.Context, Unique: models.Unique{ID: id}}).Get()
	if err != nil {
		if err == types.ErrorAccountNotSpecified || err == types.ErrorAccountNotFound {
			ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		} else {
			ctx.HandleStatus(w, r, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	ctx.HandleJson(w, r, account, http.StatusOK)
}
