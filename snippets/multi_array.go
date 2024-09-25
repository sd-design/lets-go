package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	var thisMonth = [4][7]string{{"Понедельник 2", "Вторник 3", "Среда 4", "Четверг 5", "Пятница 6", "Суббота 7", "Воскресенье 8"},
		{"Понедельник 9", "Вторник 10", "Среда 11", "Четверг 12", "Пятница 13", "Суббота 14", "Воскресенье 15"},
		{"Понедельник 16", "Вторник 17", "Среда 18", "Четверг 19", "Пятница 20", "Суббота 21", "Воскресенье 22"},
		{"Понедельник 23", "Вторник 24", "Среда 25", "Четверг 26", "Пятница 27", "Суббота 28", "Воскресенье 29"}}

	var num1 = randInt(0, 3)
	var num2 = randInt(0, 6)
	fmt.Println(thisMonth[num1][num2])
	fmt.Printf("random : %d %d\n", num1, num2)
	// fmt.Printf("random integer in range [5, 10): %d\n", randInt(4, 10))
	// fmt.Printf("random integer in range [5, 10): %d\n", randInt(4, 10))

}
