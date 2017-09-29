// Use of serverMux.HandleFunc
package main

import (
	"fmt"
	"log"
	"net/http"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use of serverMux.HandleFunc")
}

func main() {
	// Create a serveMux object
	mux := http.NewServeMux()

	// Resigter URL path with handler using ServerMux.HandleFunc
	mux.HandleFunc("/welcome", customHandler)

	log.Println("Listening...")
	// Create http server
	http.ListenAndServe(":8080", mux)
}
