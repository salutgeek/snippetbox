package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from server"))
}

// snippetView handler func
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	log.Printf("id = %d\n",id)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Display a specific snippet..."))
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

// snippetCreate handler func
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Initialize new router
	mux := http.NewServeMux()

	// Route mapping
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Start a new web server
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
