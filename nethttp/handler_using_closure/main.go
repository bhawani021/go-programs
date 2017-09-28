// Handler using closure
package main

import (
	"fmt"
	"log"
	"net/http"
)

func customHandler(body string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, body)
	})
}

func main() {
	// Create a ServerMux object
	mux := http.NewServeMux()
	// Create handler
	handler := customHandler("Closure as handler")
	// Register URL path with
	mux.Handle("/index", handler)

	log.Println("Listening...")

	// Create http server
	http.ListenAndServe(":8080", mux)
}
