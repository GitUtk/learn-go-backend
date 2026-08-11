package main

import "fmt"

func main() {
	defer fmt.Println("World") // it just remove the line and place this just above the last curly braces
	defer fmt.Println("One")
	defer fmt.Println("Two") // In the reversed order (LIFO)
	fmt.Println("Hello")
	/*
		OUTPUT:
			Hello
			4
			3
			2
			1
			0
			Two
			One
			World
	*/
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
