package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/merofuruya/search/api"
	"github.com/merofuruya/search/core/http/middleware"
)

func getRoutes() []func(r chi.Router) {
	return []func(r chi.Router){
		api.SearchRouter,
		api.TokenRouter,
	}
}

func getMiddleware() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.LoggerMiddlewareFactory(),
		// middleware.UserMiddlewareFactory(),
	}
}

func ApiRouter() chi.Router {
	router := chi.NewRouter()
	for _, middleware := range getMiddleware() {
		router.Use(middleware)
	}
	for _, route := range getRoutes() {
		router.Group(route)
	}
	return router
}
