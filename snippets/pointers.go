package main

import "fmt"

func main() {

	a := 4
	p := &a
	a += 2
	*p = *p - 1

	var answer int = 42
	var AnswerPoint *int = &answer
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(AnswerPoint)
}
