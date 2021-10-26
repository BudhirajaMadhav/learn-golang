package main

import "fmt"

const LoginToken = "lAuDaLaSsAn05458"
// 	  ↑ this Capital "L" has huge significance in Go
// It means that now this variable is public and accessible throughout its folder. 

func main() {
	var username string = "Madhav"
	fmt.Println(username)
	// %T gives a type, and "%" are known as placeholders here.
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloatVal float64 = 256.4555554875785157878475487
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of type: %T \n", smallFloatVal)
	// output is: 256.45557, for float32 -> 32 bit
	// 256.4555554875785, for float64 -> 6smallFloatVal4 bit


	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
	// output: 0

	var uninitialilsedString string
	fmt.Println(uninitialilsedString)
	fmt.Printf("Variable is of type: %T \n", uninitialilsedString)
	// default string is empty character




	// implicit way of declaring variable
	
	var website = "google.com"
	fmt.Println(website)
	// ❌ website = 3, type inference is just like kotlin.



	// no var style
	// only allowed inside a method(including main)********************
	numberOfUser := 300000.2 //(walrus)
	fmt.Println(numberOfUser)
}
