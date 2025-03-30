package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	httputil "github.com/merofuruya/search/core/http/util"
	"github.com/merofuruya/search/core/logging"
)

func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing query", http.StatusBadRequest)
		return
	}

	logger := logging.GetLogger("search")

	logger.Info().Str("query", query).Msg("searching")

	redirect := fmt.Sprintf("https://www.google.com/search?q=%s", query)

	w.Header().Set("Location", redirect)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func suggest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing query", http.StatusBadRequest)
		return
	}

	logger := logging.GetLogger("search")

	logger.Info().Str("query", query).Msg("suggesting")

	suggestions := []string{
		"hello",
		"world",
	}

	w.WriteHeader(http.StatusOK)
	httputil.JSONResponse(w, suggestions)
}

func SearchRouter(router chi.Router) {
	router.Get("/search", search)
	router.Get("/suggest", suggest)
}
