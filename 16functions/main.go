package main

import "fmt"

// We cannot write fucntion inside the functions.
func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := subtractor(3, 7)

	fmt.Println(result)

	// it printed both the values, I could have also used comma ok syntax 
	fmt.Println(proAdder(1,2,3,4,5,6,7,8,9,10))

}
                                      // to return
func subtractor(numOne int, numTwo int) int {

	return numOne - numTwo

}

// later
// anonymous functions and immediately invoked functions:
// func () {
// 	fmt.Println("anonymous function")
// }() //execute bhi ho gaya.

func greeter()  {
	fmt.Println("Namaste from golang")
}


// ---------------------------------------------------------------------

// I have no idea how many values are going to come in:
// -------------IMP-----------------
func proAdder(values ...int) (int, string) {

	total := 0

	for _, value := range values {

		total += value

	}

	return total, "go pro"
}