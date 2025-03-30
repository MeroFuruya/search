//go:build standalone

package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed all:public/dist
var content embed.FS

func StandaloneRouter(router chi.Router) {
	subFS, _ := fs.Sub(content, "public/dist")
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(subFS))))
}
