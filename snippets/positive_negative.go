package main

import "fmt"

func main() {

	/* На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное", если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.
	 */
	var input int
	fmt.Scan(&input)

	// if input > 0 {
	// 	fmt.Println("Число положительное")
	// } else if input < 0 {
	// 	fmt.Println("Число отрицательное")
	// } else {
	// 	fmt.Println("Ноль")
	// }

	switch {
	case input > 0:
		fmt.Println("Число положительное")
	case input < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Ноль")
	}
}
