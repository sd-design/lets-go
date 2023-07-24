package main

import (
	"fmt"
)

/*
Sample Input:
5
1 2 3 -1 -4
Sample Output:
3
*/

func main() {
	var workArray [100]int
	var counter int = 0
	var qty int

	fmt.Scan(&qty)
	for i := 0; i < qty; i++ {
		fmt.Scan(&workArray[i])
		if workArray[i] < 0 {
			break
		} else {
			counter++
		}
	}

	fmt.Println(counter)
}
