package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	// make is used to initialise the map	
	languages := make(map[string]string)
		
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println("JS short for:", languages["JS"])

	// deleting in map

	delete(languages, "RB")
	fmt.Println(languages)

	// iterating on maps in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

	// comma ok
	for key, _ := range languages {
		fmt.Printf("For key %v, value is v \n", key)
	}
}
