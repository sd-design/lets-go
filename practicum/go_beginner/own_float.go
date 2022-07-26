package main

import (
	"fmt"
)

func operation(num1 float64, num2 float64, digits float64) {
	result := (num1 + num2) * digits
	fmt.Println(result)
}

func main() {
	var s1, s2 float64
	var s3 float64
	fmt.Println("First number: ")
	fmt.Scanln(&s1)
	fmt.Println("Second number: ")
	fmt.Scanln(&s2)
	fmt.Println("Number of igits: ")
	fmt.Scanln(&s3)
	operation(s1, s2, s3)
}
