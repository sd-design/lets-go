package main

import "fmt"

func main() {
	m := map[int]int{
		8: 42,
		2: 6,
		4: 9,
		5: 3}
	delete(m, 2)
	fmt.Println(m[4] - m[5])
}
