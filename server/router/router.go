package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", handleRoot)
	r.Get("/{shortURL}", handleRedirect)
	r.Post("/shorten/{originalURL}", handleShortenURL)

	return r
}
