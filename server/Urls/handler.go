package Urls

import (
	"encoding/json"
	"net/http"

	"github.com/tornado-hub/url-app/server/storage"
)

func AddUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newUrl storage.ShortURL
	err := decoder.Decode(&newUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = storage.SaveUrl(newUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urls, err := storage.FindUrls()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(urls)
}

func GetLongUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	longurl := r.URL.Query().Get("longurl")
	if longurl == "" {
		http.Error(w, "Parameter 'longurl' is required", http.StatusBadRequest)
		return
	}

	originalUrl, err := storage.FindLongUrl(longurl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(originalUrl)
}
