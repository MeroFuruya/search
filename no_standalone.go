//go:build !standalone

package main

import "github.com/go-chi/chi/v5"

func StandaloneRouter(router chi.Router) {
	// no-op for non-standalone builds
}
