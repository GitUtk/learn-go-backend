package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	defer PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"
	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	defer res.Body.Close()
	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content length is:", res.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)
	//fmt.Println(content) //ByteString
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}
