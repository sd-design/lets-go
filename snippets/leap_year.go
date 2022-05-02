package main

import "fmt"

/*
Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100.
Вводится единственное число - номер года (целое, положительное, не превышает 10000).

Требуется вывести слово YES, если год является високосным и NO - в противном случае.

Sample Input:
2000
Sample Output:
YES
*/

func main() {

	var input int
	fmt.Scan(&input)
	ostatok := input % 4
	ostatok2 := input % 100
	ostatok3 := input % 400
	fmt.Println(ostatok, ostatok2, ostatok3)
	if ostatok == 0 && ostatok2 == 0 && ostatok3 == 0 {
		fmt.Println("YES")
	} else {
		//fmt.Println(ostatok2 % 4)
		if (ostatok2%4) == 0 && ostatok2 > 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
