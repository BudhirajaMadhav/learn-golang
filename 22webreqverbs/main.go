package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")

	// PerformGetRequest("http://localhost:8000")
	// PerformPostJsonRequest()
	PerformPostformRequest()
}

func PerformGetRequest(myurl string) {

	response, err := http.Get(myurl)

	if err != nil {

		panic(err)

	}

	defer response.Body.Close()
	fmt.Printf("response.Status: %v\n", response.Status)
	fmt.Printf("response.ContentLength: %v\n", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("string(content): %v\n", string(content))

	// ----------Another way of doing this

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	println(byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"
	// fake json payload

	// iska backtick aise hi aayega. dhyaan se ==========-----------================-=========-==--==-=
	requestBody := strings.NewReader(`
			{
				"name": "Madhav",
				"profession": "Programmer"
			}
		
	`)

	// ye 2nd param i.e, contentType is mentioned in the headers, which can be checked through postman etc.
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {

		panic(err)

	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder
	responseString.Write(content)

	fmt.Println(responseString.String())

}

func PerformPostformRequest() {

	const myurl = "http://localhost:8000/postform"

	// creating formdata

	data := url.Values{}

	data.Add("firstname", "madhav")
	data.Add("lastname", "budhiraja")
	data.Add("email", "madhav@google.com")

	response, err := http.PostForm(myurl, data)

	if(err != nil) {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
