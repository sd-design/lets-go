package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Println(name)
}
