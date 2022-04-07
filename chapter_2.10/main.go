package main

import (
	"flag"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
	userAgent := r.Header.Get("User-Agent")
	log.Printf("User-agent: %s ...", userAgent)
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.

	host := flag.String("host", "localhost", "network address")
	port := flag.String("port", ":4000", "HTTP port")
	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()
	addr := *host + *port
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it. Note that we're using the
	// log.Printf() function to interpolate the address with the log message.
	log.Printf("Starting server on %s - host %s", addr, *host)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
