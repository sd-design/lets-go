package main

import "fmt"

/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:
5
38 24 800 8 16
Sample Output:
40
*/

func main() {

	var qty int
	var input int
	var summ int

	fmt.Scan(&qty)
	for n := 0; n < qty; n++ {
		fmt.Scan(&input)
		if input%8 == 0 && input < 100 && input > 9 {
			summ += input
		}
	}
	fmt.Println(summ)
}
