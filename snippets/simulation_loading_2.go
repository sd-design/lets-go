package main

import (
	"fmt"
	"time"
)

func simulateLoading(percentage chan int) {

	for i := 0; i <= 100; i++ {
		percentage <- i
		time.Sleep(30 * time.Millisecond)
	}
	close(percentage)
	fmt.Println("simulation is finished")
}

func displayLoading(percentage chan int, done chan struct{}) {

	for val := range percentage {
		fmt.Println(val, "%")
	}
	fmt.Println("dislay was finished")
	close(done)
}

func main() {
	percentage := make(chan int)

	done := make(chan struct{})

	go simulateLoading(percentage)
	go displayLoading(percentage, done)
	<-done
}
