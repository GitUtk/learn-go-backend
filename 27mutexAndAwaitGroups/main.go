package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - github.com")

	var wg sync.WaitGroup
	var mu sync.Mutex

	score := []int{0}

	wg.Add(3)

	go func() {
		defer wg.Done()

		fmt.Println("One R")

		mu.Lock()
		score = append(score, 1)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()

		fmt.Println("Two R")

		mu.Lock()
		score = append(score, 2)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()

		fmt.Println("Three R")

		mu.Lock()
		score = append(score, 3)
		mu.Unlock()
	}()

	wg.Wait()

	fmt.Println(score)
}
