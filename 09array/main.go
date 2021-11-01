package main

import "fmt"

func main() {
	fmt.Println("-----------ARRAY------------")

	// Golang's specification holds our hands very much and does not allow us to do very much with it
	// so, it is not very widely used in golang.

	// Declaring the array.
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "peach"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [5]string{"potato", "cauliflower", "brinjal"}
	fmt.Println(vegList)
	
	// [potato cauliflower brinjal  ]
	// for empty value it prints a space " "

}
