package main

import "fmt"

const LoginToken string = "kjaDFasdabgsg" // Public

func main() {
	var username string = "Utkarsh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varible is of type : %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Varible is of type : %T \n", smallVal)

	var smallFloat float64 = 255.4567457674
	fmt.Println(smallFloat)
	fmt.Printf("Varible is of type : %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Varible is of type : %T \n", anotherVariable)

	//implicit type
	var website = "github.com"
	fmt.Println(website)
	// ! website = 3

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varible is of type : %T \n", LoginToken)

}
