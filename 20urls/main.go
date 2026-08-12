package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://something.com:3000/learn?course=btech&class=it"

func main() {
	fmt.Println("Handling URLs in goLang")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Print(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["class"])
	for _, val := range qparams {
		fmt.Println("Params is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "localhost",
		Path:    "/user",
		RawPath: "user=admin",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
