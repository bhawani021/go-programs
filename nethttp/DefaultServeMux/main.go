// http package provides shortcut methods for working with DefaultServeMux.
// http.Handle and http.HandleFunc
package main

import (
	"fmt"
	"log"
	"net/http"
)

func customHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use of DefaultServeMux")
}

func main() {
	// Register URL path with handler using http
	http.HandleFunc("/", customHandler)

	log.Println("Listening....")
	// Create http server
	http.ListenAndServe(":8080", nil)
}
