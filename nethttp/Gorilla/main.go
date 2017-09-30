// Design Restful API using Gorilla mux
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Note ....
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

// Stote for the Notes collection
var noteStore = make(map[string]Note)

// Variable to generate key for the collection
var id int

// PostNoteHandler ...
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// GetNoteHandler ...
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// PutNoteHandler ...
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]

	var noteToUpdate Note
	// Decode the incoming Note json
	err = json.NewDecoder(r.Body).Decode(&noteToUpdate)
	if err != nil {
		panic(err)
	}

	if note, ok := noteStore[k]; ok {
		noteToUpdate.CreatedOn = note.CreatedOn
		// Delete existing item and add the updaed item
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
		w.WriteHeader(http.StatusOK)
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteNoteHandler ...
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	// Remove from store
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
		w.WriteHeader(http.StatusAccepted)
	} else {
		log.Printf("Could not found key of Note %s to delete", k)
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
