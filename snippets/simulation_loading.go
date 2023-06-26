package main

import (
	"fmt"
	"sync"
	"time"
)

func simulateLoading(percentage chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		percentage <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(percentage)
	fmt.Println("simulation is finished")
}

func displayLoading(percentage chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range percentage {
		fmt.Println(val, "%")
	}
	fmt.Println("dislay was finished")
}

func main() {
	percentage := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go simulateLoading(percentage, wg)
	wg.Add(1)
	go displayLoading(percentage, wg)

	wg.Wait()

}
