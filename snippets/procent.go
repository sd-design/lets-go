package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	var i int = 0

	for ; x < y; i++ {
		x = x * (100 + p) / 100
		//fmt.Println(x)
	}
	fmt.Println(i)
}
