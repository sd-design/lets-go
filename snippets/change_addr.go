package main

import "fmt"

var x int = 8

func change() {
	x = 42
}
func change_ptr(ptr *int) {
	*ptr = 8
}

func main() {

	change()
	fmt.Println(x) // выведет 42

	change_ptr(&x)
	fmt.Println(x) // выведет 8
}

// func change(val int) {
// 	val = 8
// 	fmt.Println(val)
// }

// func change_ptr(ptr *int) {
// 	*ptr = 8
// }

// func main() {
// 	x := 42

// 	change(x)
// 	fmt.Println(x) // выведет 42

// 	change_ptr(&x)
// 	fmt.Println(x) // выведет 8
// 	fmt.Println(&x)
// }
