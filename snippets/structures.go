package main

import "fmt"

type Cpu struct {
	prodName  string
	Core      int
	Frequency string
}

type Pc struct {
	Cpu     Cpu
	harDisk string
}

func main() {
	var comp Pc = Pc{}
	comp.Cpu.Core = 8
	comp.Cpu.prodName = "Intel Core i7"
	comp.Cpu.Frequency = "3.4 Gz"
	fmt.Printf("Pc CPU=[%v]", comp.Cpu)

}
