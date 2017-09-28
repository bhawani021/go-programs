// Create a custom handler

// Any object can be an implementation of http.Handler if it can
// provide a method with the signature ServeHTTP(http.ResponseWriter, *http.Request)

package main

import (
	"fmt"
	"log"
	"net/http"
)

type customHandler struct {
	body string
}

func (b *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, b.body)
}

func main() {
	// Create a ServeMux object
	mux := http.NewServeMux()
	// Create a new handler
	handler := &customHandler{body: "I am a custom handler"}
	// Register URL path with handler
	mux.Handle("/home", handler)

	log.Println("Listening...")

	// Create http server
	http.ListenAndServe(":8080", mux)
}
