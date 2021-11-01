package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var numberOne int = 5
	// var numberTwo float64 = 6

	// fmt.Println("The sum is:", numberOne+int(numberTwo))

	// fmt.Println(rand.Intn(5))
	/* This gave me 1 like literally everytime.*/

	/*-------rand from crypto--------------*/

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(6))
	fmt.Println(myRandomNum)

}
