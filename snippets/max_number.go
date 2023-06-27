package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	var maxNum int = array[0] //на случай если первое число будет отрицательное или ноль
	for i, _ := range array {
		if array[i] > maxNum {
			maxNum = array[i]
		}
	}
	fmt.Println(maxNum)
}
