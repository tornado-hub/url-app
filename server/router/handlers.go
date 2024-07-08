package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tornado-hub/url-app/server/Urls"
	"github.com/tornado-hub/url-app/server/storage"
)

var shortURLs map[string]string = make(map[string]string)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func handleShortenURL(w http.ResponseWriter, r *http.Request) {
	Urls.AddUrl(w, r)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	fmt.Print(shortURL)
	originalURL, ok := storage.FindLongUrl(shortURL)
	fmt.Println(originalURL)
	if ok != nil {
		http.Error(w, "Invalid short URL", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func handleurls(w http.ResponseWriter, r *http.Request) {
	Urls.GetUrls(w, r)
}

func handleDeleteUrl(w http.ResponseWriter, r *http.Request) {
	Urls.DeleteUrl(w, r)
}
