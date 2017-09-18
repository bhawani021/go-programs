// Build a static web server
// Use FileServer

package main

import (
	"log"
	"net/http"
)

func main() {
	// Create an ServeMux object
	mux := http.NewServeMux()
	// Create a new handler
	f := http.FileServer(http.Dir("templates"))
	// Register URL path with handler
	mux.Handle("/", f)

	log.Println("Litening...")
	// Create HTTP server
	http.ListenAndServe(":8080", mux)
}
