package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	originalURL := chi.URLParam(r, "originalURL")

	if originalURL == "" {

		t, err := template.ParseFiles("form.html")
		if err != nil {
			fmt.Fprintf(w, "Error parsing template: %v", err)
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

func redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	originalURL, ok := shortURLs[shortURL]
	if !ok {
		fmt.Fprintf(w, "Invalid short URL")
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		randomString := generateRandom()
		w.Write([]byte(randomString))
	})
	r.Get("/{shortURL}", redirect)
	//r.HandleFunc("/shorten/{originalURL}", shortenUrl).Methods("POST")
	http.ListenAndServe(":3000", r)

	randomString := generateRandom()
	fmt.Println(randomString)
}
