package main

import (
	"bytes"
)

func CBytes(args []string) string {
	buffer := bytes.Buffer{}
	for _, arg := range args {
		buffer.WriteString(arg)
	}
	return buffer.String()
}

var params = []string{"Who are", "you man", "? I'm Android", " by Golang"}

func main() {
	println(CBytes(params))
}
