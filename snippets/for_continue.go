package main

import "fmt"

func main() {
	sum, limit := 0, 100
	for i := 0; true; i++ {
		if i%2 != 0 {
			fmt.Println("нечетное")
			continue // переход к следующему числу, так как i — нечётное
		}

		if sum+i > limit {
			break // выход из цикла, так как сумма превысит заданный предел
		}

		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
