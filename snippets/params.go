package main

import (
	"fmt"
	"os"
)

func main() {

	// The first argument
	// is always program name
	myProgramName := os.Args[0]
	toGetAllArgs := os.Args[1:]

	// it will display
	// the program name
	fmt.Println(myProgramName)
	fmt.Println(toGetAllArgs)
}
