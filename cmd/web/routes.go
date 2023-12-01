package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Initialize new router
	mux := http.NewServeMux()
	// Route mapping
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.Handle(
		"/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))),
	)

	return mux
}
