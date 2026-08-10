package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)  //actual memory address
	fmt.Println("Value of actual pointer is ", *ptr) // actual value inside the pointer

	*ptr = *ptr + 2
	fmt.Println("New value is ", myNumber) // actual value inside the pointer

}
