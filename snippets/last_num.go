package main

import (
	"fmt"
)

func main() {
	var num string
	var lenght int
	fmt.Scan(&num)
	lenght = len(num)
	fmt.Println(string(num[lenght-2]))
}
