package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Header().Set("Cache-Control", "public, max-age=31536000")
	// w.Write([]byte("Hello from Snippetbox"))
	/*
		Supress the Date header
	*/
	// w.Header()["Date"] = nil
	/*
		Set the Secure-Token header
	*/
	w.Header().Set("Secure-Token", "mySecureToken2021")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}
func getHeader(w http.ResponseWriter, r *http.Request) {
	/*
		Example how get header and check it
		My-header : fhskjdh5623lkhlk4234980&3234j_9432849
	*/
	checkHeader := "fhskjdh5623lkhlk4234980&3234j_9432849"
	idHeader := r.Header.Get("My-header")
	if idHeader != checkHeader {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "You can see header in developer tools of browser -  %s", idHeader)
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not. Note that
	// http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", showSnippet)
	mux.HandleFunc("/header/get/", getHeader)
	mux.HandleFunc("/snippet/create/", createSnippet)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
