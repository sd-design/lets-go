package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	Username  string
	Password  string `json:"-"`
	IsAdmin   bool
	CreatedAt time.Time
}

func main() {
	addr := ":5000"
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", echoHandler)
	log.Printf("Starting server on %s", addr)
	http.ListenAndServe(addr, mux)

}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	if r.Method != http.MethodPost {
		// Use the Header().Set() method to add an 'Allow: POST' header

		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	json.NewDecoder(r.Body).Decode(&user)

	user.CreatedAt = time.Now().Local()

	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}
