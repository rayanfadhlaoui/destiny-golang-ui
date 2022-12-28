package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rayanfadhlaoui/destiny-golang-ui/pkg/handlers"
	"net/http"
)

func routes(r *handlers.Repository) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Get("/", r.Home)
	mux.Get("/about", r.About)
	return mux
}
