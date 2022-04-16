package main

import "fmt"

func main() {

	a := 4
	p := &a
	a += 2
	*p = *p - 1
	fmt.Println(a)
}
