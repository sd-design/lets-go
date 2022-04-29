package main

import "fmt"

func main() {

	/*
		На вход дается натуральное число, не превосходящее 10000.
		Формат выходных данных
		Выведите одно целое число - первую цифру заданного числа.
	*/
	var input int
	var res int
	fmt.Scan(&input)
	if (input / 10000) != 0 {
		res = input / 10000
	} else if (input / 1000) != 0 {
		res = input / 1000
	} else if (input / 100) != 0 {
		res = input / 100
	} else if (input / 10) != 0 {
		res = input / 10
	} else {
		res = input
	}

	fmt.Println(res)

	//Variation by switch
	switch {
	case input/10000 >= 1:
		fmt.Println(input / 10000)
	case input/1000 >= 1:
		fmt.Println(input / 1000)
	case input/100 >= 1:
		fmt.Println(input / 100)
	case input/10 >= 1:
		fmt.Println(input / 10)
	default:
		fmt.Println(input)
	}
}
