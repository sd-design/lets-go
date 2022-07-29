package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a string
	a = "абцд"
	fmt.Println(len(a))                    // 8
	fmt.Println(utf8.RuneCountInString(a)) //4
}
