package main

import (
	"fmt"
	_ "net/http"

	_ "github.com/tornado-hub/url-app/server/router"
	"github.com/tornado-hub/url-app/server/storage"
	//"github.com/tornado-hub/url-app/server/Urls"
)

func main() {
	storage.InitDB("shorturl.db")
	fmt.Print(storage.FindUrls())
	//r := router.NewRouter()

	fmt.Println("Server listening on port 3000...")
	//http.ListenAndServe(":8000", r)
}
