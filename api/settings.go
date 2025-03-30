package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	httputil "github.com/merofuruya/search/core/http/util"
)

func token(w http.ResponseWriter, r *http.Request) {
	httputil.JSONResponse(w, map[string]string{
		"token": "This is a token",
	})
}

func TokenRouter(router chi.Router) {
	router.Get("/token", token)
}
