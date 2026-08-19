package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang - github.com/gitUtk")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Sender goroutine
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println("Sending 5...")
		ch <- 5
		fmt.Println("Sent 5")
	}(myCh, wg)

	// Receiver goroutine
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		value := <-ch
		fmt.Println("Received:", value)
	}(myCh, wg)

	wg.Wait()

	fmt.Println("Main function completed")
}
