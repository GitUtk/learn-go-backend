package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	utkarsh := User{"Utkarsh", "test@test.com", true, 20}
	fmt.Println(utkarsh)
	fmt.Printf("%+v\n", utkarsh) // mpre detailed version
	fmt.Printf("Age is %v and email is %v", utkarsh.Age, utkarsh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
