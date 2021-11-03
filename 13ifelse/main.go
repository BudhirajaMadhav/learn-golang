package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10
	
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	// weird but common syntax in go. sometimes a value might come thru a web request and we just want to initialise it on the go, so GO.
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")	
	} else {
		fmt.Println("Number is NOT less than 10")
	}

	// if err != nil {}

}
