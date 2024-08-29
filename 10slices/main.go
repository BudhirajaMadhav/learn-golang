package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// if we're using this syntax then we need to initialise the slice as well.
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	// Output: Type of fruitList is []string

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	// sliced like python. 1st index se last tak.
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 124
	highScores[1] = 568
	highScores[2] = 74

	fmt.Println(highScores)

	highScores = append(highScores, 874)

	fmt.Println(highScores)

	// Output: [124 568 74 0 874] -> jo make memory wala hai, matlab jo allocated wala hai wo index se hi edit hoga. Like, in this case wo 3 index 0 hi rehgya.

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	/*
		How to remove a value from slices based on index.
	*/

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2

	// this is known as ellipsis
	courses = append(courses[:index], courses[index + 1:]...)
	fmt.Println(courses)

}
