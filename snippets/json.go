package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name      string
	Age       int
	IsBlocked bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	newUser := User{
		Name:      "Alex",
		Age:       43,
		IsBlocked: false,
	}
	jsonLine, _ := json.Marshal(newUser)
	fmt.Println(string(jsonLine))
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(jsonLine)))
	// w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
