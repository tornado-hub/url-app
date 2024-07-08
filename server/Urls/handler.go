package Urls

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/tornado-hub/url-app/server/storage"
)

const chars string = "1234567890abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVQXYZ"

func generateRandom() string {
	var shortString string
	for i := 0; i < 6; i++ {
		x := rand.Intn(len(chars))
		shortString += string(chars[x])
	}
	return shortString
}

func AddUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newUrl storage.ShortURL
	err := decoder.Decode(&newUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var shortURL string
	for {
		shortURL = generateRandom()
		// Check if the generated short URL already exists
		_, err := storage.FindLongUrl(shortURL)
		if err != nil && err.Error() == "sql: no rows in result set" {
			// Short URL does not exist, it's unique
			break
		}
	}

	newUrl.ShortURL = shortURL

	err = storage.SaveUrl(newUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]string{
		"short_url":    shortURL,
		"original_url": newUrl.OriginalURL,
	}
	json.NewEncoder(w).Encode(response)
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

func GetLongUrl(w string) {

	originalUrl, _ := storage.FindLongUrl(w)

	fmt.Println(originalUrl)
}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req struct {
		ShortURL string `json:"short_url"`
	}
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = storage.DeleteUrl(req.ShortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func GetLongUrl(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	longurl := r.URL.Query().Get("longurl")
// 	if longurl == "" {
// 		http.Error(w, "Parameter 'longurl' is required", http.StatusBadRequest)
// 		return
// 	}

// 	originalUrl, err := storage.FindLongUrl(longurl)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(originalUrl)
// }
