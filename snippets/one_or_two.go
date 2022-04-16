package main

import "fmt"

/*
Вам нужно написать функцию one_or_two().  На вход подается два целых числа и строка. Строка может иметь значения one, two или любой другой текст.

Возвращать из функции вам нужно два значения. Если строка равна one, нужно вернуть первое число и саму строку. Если строка равна two, нужно вернуть второе число и саму строку. Если строка другая - нужно вернуть 0 и саму строку.

Sample Input: 2 5 two
Sample Output: 5 two
*/
func one_or_two(a int, b int, text string) (int, string) {
	if text == "one" {
		return a, text
	}
	if text == "two" {
		return b, text
	}
	return 0, text
}

func main() {

	num, text := one_or_two(2, 5, "tree")

	fmt.Println(num, text)
}
