package main

import "fmt"

func main() {
	list := [11]string{"Ноль", "Один", "Два", "Три", "Четыре", "Пять", "Шесть", "Семь", "Восемь", "Девять", "Десять"}
	var num1, num2, num3 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)
	fmt.Println(list[num1])

	fmt.Println(list[num2])

	fmt.Println(list[num3])
}
