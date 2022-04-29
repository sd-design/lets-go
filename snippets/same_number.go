package main

import "fmt"

func main() {

	/*
		Выведите "YES", если все цифры числа различны, в противном случае - "NO".
	*/
	var input int
	fmt.Scan(&input)
	num1 := input / 100
	num2 := (input - (num1 * 100)) / 10
	num3 := (input - (num1 * 100)) - (num2 * 10)

	if num1 != num2 && num1 != num3 && num2 != num3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println(num1, num2, num3)
}
