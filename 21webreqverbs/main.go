package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// defer PerformGetRequest()
	// defer PerformPostJsonRequest()
	defer PerformPostFormRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"course":"Let's go with golang",
			"price":0,
			"platform":"github.com"
		}
	`)
	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	//formdata
	data := url.Values{}
	data.Add("Firstname", "Utkarsh")
	data.Add("LastName", "Yadav")
	data.Add("email", "utkarsh@go.dev")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))

	defer res.Body.Close()
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
