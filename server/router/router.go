package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func NewRouter() http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	r := chi.NewRouter()

	// Middleware
	r.Use(corsHandler.Handler)

	// Routes
	r.Get("/", handleRoot)
	r.Get("/{shortURL}", handleRedirect)
	r.Post("/shorten", handleShortenURL)
	r.Get("/urls", handleurls)
	r.Delete("/delete", handleDeleteUrl)

	return r
}
