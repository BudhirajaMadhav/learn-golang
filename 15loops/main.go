package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Tuesdsay", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// go does not have ++i -----------------------------------------------------------
	// for i := 0; i < len(days); i++ {
		// fmt.Println(days[i])
	// }

	// this is not foreach, i is index here and its scope is limited to this loop.
	for i := range days{
		fmt.Println(days[i])
	}
	
	// for each kind of loop.

	// comma ok
	for index, day := range days{
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	fmt.Println("------------==================-==--=-=-========")

	rougeValue := 1

	for rougeValue < 10 {

		// goto ---------------------------------
		if rougeValue == 3 {
			goto gomc
		}

		if rougeValue == 5 {
			break
		}

		fmt.Println("Vaule is:", rougeValue)
		rougeValue++
	}

	fmt.Println("===========================")


	// label ----------------------------------------
	gomc:
		fmt.Println("go mc hai")

	



}
