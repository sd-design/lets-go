package main

import "fmt"

func codes(hello string) {
	for _, v := range hello {
		fmt.Println(v)
	}
}

func main() {
	hello := "Hello Go . It's a harm"
	codes(hello)
	fmt.Println("Hello Go . It's a harm"[0])
	fmt.Println(string("Hello Go"[0:3]))
	fmt.Println(len("Hello World"))
}
