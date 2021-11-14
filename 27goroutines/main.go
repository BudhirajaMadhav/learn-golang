package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{"test"}
var mut sync.Mutex

func main() {

	// main method by itself does not wait for the thread to come back.
	// 1. sleep time = not a good solution.
	// go greeter("hello")
	// greeter("world")


	websiteList := []string{
		"https://gooogle.com",
		"https://fb.com",
		"https://instagram.com",
		"https://twitter.com",
		"https://go.dev",
		
	}

	for _, site := range websiteList {
		go getStatusCode(site)
		wg.Add(1)
		// for this, add, done and wait combo to work:
		// add should fire before done is fired.
		// wait will keep the func from ending until the goroutine count reaches 0
	}

	wg.Wait() //It waits here, at this line.
	fmt.Println(signals)

}

// func greeter(s string) {

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}

// }

func getStatusCode(endpoint string) {

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint", err)
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

	signals = append(signals, endpoint)

	wg.Done()
	
}