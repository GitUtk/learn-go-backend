package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Introduction to slicing on the array")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruiteList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 777) // Memory re allocation

	fmt.Println(highScores)

	sort.Ints(highScores) // Sorting the slices
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // true or false based on if sorted

	// How to remove a value from slices based on index
	var courses = []string{"python", "java", "go", "cpp", "reactjs"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
