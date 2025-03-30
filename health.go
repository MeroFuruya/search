package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	httputil "github.com/merofuruya/search/core/http/util"
)

type HealthMessage struct {
	Status string `json:"status"`
}

func healthHandler(writer http.ResponseWriter, request *http.Request) {
	msg := HealthMessage{Status: "ok"}
	httputil.JSONResponse(writer, msg)
}

func HealthRouter(router chi.Router) {
	router.Get("/health", healthHandler)
}
