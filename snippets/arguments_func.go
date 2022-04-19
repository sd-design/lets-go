package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func arguments(words ...string) {
	for _, v := range words {
		fmt.Println(v)
	}
}

func main() {
	sum(2, 4, 6) // выведет 12
	sum(42, 8)   // выведет 50
	sum(1, 2, 3, 4, 5, 6)
	arguments("hello", "this", "is", "arguments")
}
