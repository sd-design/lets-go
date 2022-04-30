package main

import (
	"fmt"
	"strconv"
)

/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
На вход подается номер билета - одно шестизначное  число.
Выведите "YES", если билет счастливый, в противном случае - "NO".
Sample Input:

613244
Sample Output:
YES
*/

func main() {
	// var input int
	// fmt.Scan(&input)
	// fmt.Println(input)
	// num1 := input / 100000
	// num2 := (input - (num1 * 100000)) / 10000
	// fmt.Println(num1, num2)
	const (
		yes = "YES"
		no  = "NO"
	)
	var input string
	fmt.Scan(&input)
	num1, _ := strconv.Atoi(input[0:1])
	num2, _ := strconv.Atoi(input[1:2])
	num3, _ := strconv.Atoi(input[2:3])
	num4, _ := strconv.Atoi(input[3:4])
	num5, _ := strconv.Atoi(input[4:5])
	num6, _ := strconv.Atoi(input[5:6])

	if (num1 + num2 + num3) == (num4 + num5 + num6) {
		fmt.Println(yes)
	} else {
		fmt.Println(no)
	}

}
