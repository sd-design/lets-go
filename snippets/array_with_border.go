package main

import (
	"fmt"
)

func main() {
	var qty int
	var a int
	var array [100]int

	fmt.Scanln(&qty)
	for i := 0; i < qty; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	for i := 0; i < qty; i++ {
		fmt.Println(array[i])
	}

}
