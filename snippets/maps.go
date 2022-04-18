package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	//also initialize of map
	myMap2 := map[string]int{
		"James": 42,
		"Amy":   24}

	myMap["hello"] = 1

	fmt.Println(myMap["hello"])
	fmt.Println(myMap2["Amy"])

}
