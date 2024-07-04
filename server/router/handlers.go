package router

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var shortURLs map[string]string = make(map[string]string)

const chars string = "1234567890abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVQXYZ"

func generateRandom() string {
	rand.Seed(time.Now().UnixNano())
	var shortString string
	for i := 0; i < 6; i++ {
		x := rand.Intn(len(chars))
		shortString += string(chars[x])
	}
	return shortString
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	randomString := generateRandom()
	w.Write([]byte(randomString))
}

func handleShortenURL(w http.ResponseWriter, r *http.Request) {
	originalURL := chi.URLParam(r, "originalURL")

	if originalURL == "" {
		t, err := template.ParseFiles("form.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}

	shortURL := generateRandom()
	for shortURLs[shortURL] != "" {
		shortURL = generateRandom()
	}
	shortURLs[shortURL] = originalURL

	fmt.Fprintf(w, "Your shortened URL: %s/%s", "http://localhost:3000", shortURL)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	originalURL, ok := shortURLs[shortURL]
	if !ok {
		http.Error(w, "Invalid short URL", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
