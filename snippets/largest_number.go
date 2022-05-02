package main

import "fmt"

func main() {
	var n int
	var largest int
	var qty int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if largest == 0 {
			largest = n
			qty = 1
		} else if n > largest {
			largest = n
			qty = 1
		} else if n == largest {
			qty++
		}
	}
	fmt.Println(qty)
}
