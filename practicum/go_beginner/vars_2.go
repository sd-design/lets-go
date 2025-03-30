package main

import "fmt"

type UserID int

func main() {
	idx := 1

	var uid UserID = 42

	myID := UserID(idx)

	fmt.Println(uid, myID)

}
