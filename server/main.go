package main

import (
	"fmt"
	"net/http"

	"github.com/tornado-hub/url-app/server/router"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Server listening on port 3000...")
	http.ListenAndServe(":3000", r)
}
