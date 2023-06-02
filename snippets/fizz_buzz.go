package main

import "fmt"

func main() {
	limit := 100
	for i := 1; true; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Printf("FizzBuzz: %d\n", i)
			continue
		}
		if i%3 == 0 {
			fmt.Printf("Fizz: %d\n", i)
			continue
		}
		if i%5 == 0 {
			fmt.Printf("Buzz: %d\n", i)
			continue
		}

		if i > limit {
			break // выход из цикла, так как сумма превысит заданный предел
		}

	}
}
