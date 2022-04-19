package main

import "fmt"

func main() {
	a := [5]int{0, 2, 4, 6, 8}
	var s []int = a[2:4]

	s[0] = 8
	fmt.Println(s[1])
	x := []int{1, 2, 3, 4, 5, 6}
	var xxl []int = x[4:] //Take elements of slices from 4-th element
	fmt.Println(xxl)
}
