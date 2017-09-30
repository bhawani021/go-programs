// Use http.Server struct
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use of http.Server struct")
}

func main() {
	// Resiger URL path with handler
	http.HandleFunc("/", customHandler)

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Listening....")
	server.ListenAndServe()
}
