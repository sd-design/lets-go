package main

import (
	"fmt" // New import
	"log"
	"net/http"
	"strconv" // New import
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	hello := r.URL.Query().Get("hello")
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d ...", id)
	fmt.Fprintf(w, "Display hello param: %s ...", hello)
}

func main() {
	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
