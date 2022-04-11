package main

import (
	"fmt"
	"os"
)

func main() {

	command := os.Args[1]
	myFilename := os.Args[2]

	if command == "create" {
		file, err := os.Create(myFilename)
		if err != nil {
			fmt.Printf("Error of creating file [%s]", err.Error())
		}
		fmt.Printf("File %s was created\n", file.Name())
	}

}
