package router

import (
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tornado-hub/url-app/server/Urls"
)

var shortURLs map[string]string = make(map[string]string)

const chars string = "1234567890abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVQXYZ"

func generateRandom() string {
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
	Urls.AddUrl(w, r)
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

func handleurls(w http.ResponseWriter, r *http.Request) {
	Urls.GetUrls(w, r)
}
