package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	// Output:
	// Hello
	// Two
	// One
	// World


	// defer puts the statement at the end of the function just before the "}"

	myDefer()
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}