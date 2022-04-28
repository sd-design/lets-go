package main

import "fmt"

func main() {

	type ByteSize float64

	const (
		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fmt.Println(KB)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6

	const (
		a  = iota + 1 //iota = 0
		_             //iota = 1
		b             //iota = 2
		c             //iota = 3
		d  = c + 2    //iota = 4
		t             //iota = 5
		i             //iota = 6
		i2 = iota + 2 //iota = 7
	)
	fmt.Println(a, b, c, d, t, i, i2)

	var double bool = a != 8 // false
	fmt.Println(double)

}
