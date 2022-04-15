package main

import "fmt"

/*
Напишите функцию double_m(), которая должна принимать на вход два целых числа a и b и возвращать сумму квадратов чисел от a до b (включительно).
*/
func double_m(a, b int) int {
	var res int = 0

	for i := a; b >= i; i++ {
		res += (i * i)
	}
	return res
}

func main() {
	response := double_m(2, 6)
	fmt.Println(response)
}
