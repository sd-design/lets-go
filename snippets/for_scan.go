package main

import "fmt"

func main() {
	var n int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}
}
