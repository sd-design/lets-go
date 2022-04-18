package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func (c *Contact) drive(speed int) {
	c.age += speed
}

func main() {
	x := Contact{"Джеймс", 42}
	x.drive(8)
	fmt.Println(x.age)
}
