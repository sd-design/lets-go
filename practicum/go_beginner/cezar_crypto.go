package main

import (
	"fmt"
)

func main() {
	var rus string
	rus = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	fmt.Println(len(rus))
	fmt.Println(string(rus[64:66]))
}
