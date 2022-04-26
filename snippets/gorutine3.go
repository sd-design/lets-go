package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ch2 := make(chan string)
	go say("hello", ch)
	go say("world", ch2)
	fmt.Println(<-ch)
	fmt.Println(<-ch2)
}

func say(word string, ch chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println(word)
	ch <- "exit"
}
