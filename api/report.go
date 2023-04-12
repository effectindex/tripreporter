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

func SetupReportEndpoints(v1 *mux.Router) {
	a1 := v1.Methods(http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodPost).Subrouter()
	a1.Use(AuthMiddleware())

	a1.HandleFunc("/report", ReportPost).Methods(http.MethodPost)
	a1.HandleFunc("/report/{id}", ReportGet).Methods(http.MethodGet)
	a1.HandleFunc("/report/{id}", ReportPatch).Methods(http.MethodPatch)
	a1.HandleFunc("/report/{id}", ReportDelete).Methods(http.MethodDelete)
}

// ReportPost path is /api/v1/report
func ReportPost(w http.ResponseWriter, r *http.Request) {
	ctxVal, ok := ctx.GetCtxValOrHandle(w, r)
	if !ok {
		return
	}

	report, err := (&models.Report{Context: ctx.Context, Account: ctxVal.Account}).FromBody(r)
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	report, err = report.Post()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.HandleJson(w, r, report, http.StatusCreated)
}

// ReportGet path is /api/v1/report/{id}
func ReportGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, _ := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.Handle(w, r, MsgNotFound)
		return
	}

	report, err := (&models.Report{Context: ctx.Context, Unique: models.Unique{ID: id}}).Get()
	if err != nil {
		ctx.HandleStatus(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.HandleJson(w, r, report, http.StatusOK)
}

// ReportPatch path is /api/v1/report/{id}
func ReportPatch(w http.ResponseWriter, r *http.Request) {
	ctx.Handle(w, r, MsgNotImplemented)
}

// ReportDelete path is /api/v1/report/{id}
func ReportDelete(w http.ResponseWriter, r *http.Request) {
	ctx.Handle(w, r, MsgNotImplemented)
}
