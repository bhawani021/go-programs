// Use functions as handlers
// Use http.HandlerFunc type to server as an HTTP handler.
// http.HandlerFunc works as an adapter that allows a nomal function to work as HTTP handler
// http.HandlerFunc has a built-in method ServerHTTP(http.ResponseWriter, *http.Request)

package main

import (
	"fmt"
	"log"
	"net/http"
)

func customHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Example of function as handler")
}

func main() {
	// Create a ServerMux object
	mux := http.NewServeMux()
	// Convert customHandler function to a HandlerFunc type
	handler := http.HandlerFunc(customHander)
	// Register URL path with handler
	mux.Handle("/home", handler)

	log.Println("Listening...")

	// Create http server
	http.ListenAndServe(":8080", mux)
}
