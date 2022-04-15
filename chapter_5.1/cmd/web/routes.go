package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.HandleFunc("/api/snippet/", app.apiShowSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
