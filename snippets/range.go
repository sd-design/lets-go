package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[1] = 1
	a[2] = 3
	a[4] = 4

	for i, v := range a {
		fmt.Println(i, v)
	}
	//Only values of array
	for i := range a {
		fmt.Println(a[i])
	}
	//light version of only values of array
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("%c \n", c)
}
