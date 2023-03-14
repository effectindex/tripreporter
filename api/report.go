package api

import (
	"net/http"

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
	ctx.HandleMessage(MsgNotImplemented)
}

// ReportGet path is /api/v1/report/{id}
func ReportGet(w http.ResponseWriter, r *http.Request) {
	ctx.HandleMessage(MsgNotImplemented)
}

// ReportPatch path is /api/v1/report/{id}
func ReportPatch(w http.ResponseWriter, r *http.Request) {
	ctx.HandleMessage(MsgNotImplemented)
}

// ReportDelete path is /api/v1/report/{id}
func ReportDelete(w http.ResponseWriter, r *http.Request) {
	ctx.HandleMessage(MsgNotImplemented)
}
