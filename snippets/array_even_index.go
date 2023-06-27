package main

import (
	"fmt"
)

/*
1.12 Массивы и срезы
Sample Input:
6
4 5 3 4 2 3

Sample Output:
4 3 2
*/

func main() {
	var qty int
	var a int
	var array [100]int

	fmt.Scanln(&qty)

	for i := 0; i < qty; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	for i := 0; i < qty; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", array[i])
		}
	}

}
