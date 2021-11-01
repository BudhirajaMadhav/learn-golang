package main

import "fmt"

func main() {
	fmt.Println("----------POINTERS-----------")

	// To create a pointer
	// var ptr *int
	// fmt.Println(ptr) //prints <nil>, since ptr does point to nothing.
	// -----------We've just created a pointer which points to no memory address for now.
	// So, for that we need to create a variable jisse memory assign hogi fir jiska address pointer me daalenge

	var number = 56
	ptr := &number // & se uska memory reference milgya

	// * se deference
	*ptr = number + 27

	fmt.Println(number)
	// prints 83.
}
