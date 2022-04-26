package main

import "fmt"

func main() {
	var num1 float64 = 1.01
	var num2 float64 = 0.99
	res := num1 - num2

	s := fmt.Sprintf("%.2f", res)
	fmt.Println(s)
}
