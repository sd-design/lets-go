package main

import "fmt"

func countWin(items ...string) int {
	wins := 0
	for _, v := range items {
		if v == "w" {
			wins += 3
		}
		if v == "d" {
			wins += 1
		}
	}
	return wins
}
func main() {
	var input string

	fmt.Scanln(&input)
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	results = append(results, input)
	fmt.Println(countWin(results...))

}
