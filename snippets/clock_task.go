package main

import "fmt"

func main() {
	var degree int
	fmt.Scan(&degree)
	time := degree * 2
	hours := time / 60
	minutes := time - (hours * 60)
	// fmt.Println(time)
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
