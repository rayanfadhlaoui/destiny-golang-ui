package main

import (
	"destiny-go-ui/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(r *handlers.Repository) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer)
	mux.Get("/", r.Home)
	mux.Get("/about", r.About)
	return mux
}
