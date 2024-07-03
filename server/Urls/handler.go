// code taken from github.com/katzien/go-structure-examples

package Urls

import (
	"encoding/json"
	"net/http"
)

func AddUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newUrl shortURL
	err := decoder.Decode(&newUrl)
	if err != nil {
		return
	}
	storage.DB.SaveUrl(newUrl)

}
func GetUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ManyUrl := storage.DB.FindUrls()
	json.NewEncoder(w).Encode(ManyUrl)
}
func GetLongUrl(longurl string) string {
	LongUrl, _ := storage.DB.FindLongUrl(shortURL{OriginalURL: longurl})
	if len(LongUrl) == 1 {
		return LongUrl[0]
	}
}
