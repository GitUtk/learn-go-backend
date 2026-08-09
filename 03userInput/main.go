package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) // For taking input we use bufio
	fmt.Println("Enter the rating for our pizza:")

	// comma ok  || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T : ", input)

}
