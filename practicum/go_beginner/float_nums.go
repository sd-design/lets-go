package main

import (
	"fmt"
)

func main() {

	var x float64 = 1.01
	var i float64 = 0.01

	for x < 1.4 {
		fmt.Println(x)
		x += i
	}

}
