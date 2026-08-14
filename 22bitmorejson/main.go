package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // now this filed will not be reflected
	Tags     []string `json:"tags,omitempty"` // now will not show nil data
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{
			"ReactJS",
			299,
			"Github",
			"abc123",
			[]string{"web-dev", "js"},
		},
		{
			"MERN",
			199,
			"Github",
			"bcd123",
			[]string{"Full stack", "js"},
		},
		{
			"Angular",
			299,
			"Github",
			"utk123",
			nil,
		},
	}

	//package this data as JSON data

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJS",
                "Price": 299,
                "website": "Github",
                "tags": ["web-dev","js"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type of data is %T\n", k, v, v)
	}
}
