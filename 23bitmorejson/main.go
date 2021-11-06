package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename` // aliases when throwing this init json.
	Price    int
	Platform string   `json:"website`
	Password string   `json:"-"`               // I don't want this field to be reflected to whoever is consuming this API
	Tags     []string `json:"tags,omitempty"` // If that field is null or nil, so just don't throw the field at all. NO SPACE BETWEEN COMMA AND OMITEMPTY.
}

func main() {
	fmt.Println("JSON in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"reactjs bootcamp", 299, "youthoob", "abc123", []string{"web-dev", "js", "react"}},
		{"MERN bootcamp", 594, "youthoob", "bct123", []string{"full-stack", "js"}},
		{"Angular bootcamp", 891, "youthoob", "mad234", nil},
	}

	// package this data into json data.

	finalJson, err := json.Marshal(lcoCourses)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

	// hard to consume,
	// password visible,
	// last one is null, not nil.

	fmt.Println("")

	finalJson, err = json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	// we see that data comes from the web, it comes in the byte format, and till now 
	// we've only converted that byte data into string, but this time we're gonna 
	// convert that data into json.
	jsonDataFromWeb :=[]byte(`
	{
		"Name": "reactjs bootcamp",
		"Price": 299,
		"Platform": "youthoob",
		"tags": [
				"web-dev","js","react"]
	}
	`)

	var mygocourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON data was valid")
		json.Unmarshal(jsonDataFromWeb, &mygocourse)
		// %#v is go syntax representation of the value. 
		fmt.Printf("%#v\n", mygocourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}


	// some cases where you just want to add data to key value.

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("The key is %v, value is %v and the type of the value is %T\n", key, value, value)
	}

}