package main

import (
	"fmt"
)

func main() {

	myBool := false
	myString := "Hello people"

	if !myBool {
		var x bool
		fmt.Println(x)
		fmt.Println(myString)
		a := 11
		b := 6
		fmt.Println(a % b)
	} else {
		fmt.Println("Something doesn't go good")
	}

}
