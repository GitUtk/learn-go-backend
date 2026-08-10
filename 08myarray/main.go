package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	// fruitList[2] = "tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is:", len(fruitList)) // 4

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Vegy list is: ", vegList)

}
