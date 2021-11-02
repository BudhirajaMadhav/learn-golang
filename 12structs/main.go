package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang, No super or parent.

	madhav := User{"Madhav", "madhav@google.com", true, 18}
	fmt.Println(madhav)

	fmt.Printf("Details of madhav are: %+v \n", madhav)
	fmt.Printf("Name is: %v\n", madhav.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
