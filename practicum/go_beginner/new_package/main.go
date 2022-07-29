package main

import (
	"fmt"

	"example.com/mysnippets/pkg/contacts"
)

func main() {
	contacts.SetSupport("Служба поддержки")
	fmt.Println(contacts.GetContact())
	fmt.Println("Email:", contacts.Email)
}
