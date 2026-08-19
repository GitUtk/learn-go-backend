package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer

var mut sync.Mutex // pointer

func main() {
	// go greeter("Hello")
	// greeter("world")
	websiteList := []string{
		"https://github.com",
		"https://go.dev",
		"https://vercel.com",
		"https://youtube.com",
		"https://google.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d stats code for %s\n", res.StatusCode, endpoint)
	}
}
