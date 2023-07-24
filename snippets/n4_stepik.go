package main

import "fmt"

func n4() {

	var workArray [10]uint8

	// fmt.Scan(&num1)

	// fmt.Println(num1)
	for i, _ := range workArray {
		fmt.Scan(&workArray[i])
		fmt.Println(i)
	}
	fmt.Println(workArray)

}

func main() {

	n4()

}
