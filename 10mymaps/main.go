package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages: ", language)
	fmt.Println("JS shorts for: ", language["JS"])

	delete(language, "RB")
	fmt.Println("List of all languages: ", language)

	//loops
	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
