package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in go lang")
	greeter()
	// 	func greeterTwo(){
	// 	fmt.Println("Another method")		// WE CAN'T DO THIS
	// }
	greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	// fmt.Println("Total sum is : ", proAdder(4, 54, 3, 53, 53, 6, 4, 4))
	sum, message := proAdder(345, 345, 345, 34, 5345, 4)
	fmt.Println("My message is ", message, "and sum is ", sum)
}

func greeter() {
	fmt.Println("Hello from golang")
}

func greeterTwo() {
	fmt.Println("Another method")
}

func adder(a int, b int) int { // Also define the data type we are returning also known as signature
	return a + b

}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "It works"
}
