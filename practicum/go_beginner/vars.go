package main

import "fmt"

func main() {

	myVar := 30
	//myVar = 31
	myVar++

	var myBool bool

	var myByte byte = '\x26'

	var myRune1 rune = '0'

	var myRune2 rune = '1'

	fmt.Println(myVar)
	fmt.Println(myBool)
	fmt.Println(myByte)
	fmt.Println("RUNEs:\n\n", myRune1)
	fmt.Println(myRune2)

	fmt.Println("\\U0031")

}
