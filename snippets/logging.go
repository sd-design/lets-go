package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("Starting logging")
}
