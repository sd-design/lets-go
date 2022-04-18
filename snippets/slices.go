package main

import "fmt"

func main() {
	a := [5]int{0, 2, 4, 6, 8}
	var s []int = a[2:4]

	s[0] = 8
	fmt.Println(s[1])
}
