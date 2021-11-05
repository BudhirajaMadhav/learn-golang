package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=643hfufasd"

// basic url syntax

func main() {

	fmt.Println("Welcome to handling urls in golang")

	// parsing url: extracting

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Port())

	// in js this is called params/parameters
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	// This is a map.
	fmt.Printf("The type of query params is: %T \n", qparams)
	fmt.Println(qparams)

	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "google.com",
		RawQuery: "q=hello",
		Path:     "/search",
	}

	anotherURL := partsOfURL.String()

	fmt.Println(anotherURL)

	resp, _ := http.Get(anotherURL)
	databytes, _ := ioutil.ReadAll(resp.Body)
	content := string(databytes)
	fmt.Println(content)
}
