package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content:= "this needs to go in a file"

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	
	checkNilErr(err)
	
	length, err := io.WriteString(file, content)
	
	checkNilErr(err)
	
	fmt.Println("Length is:", length)
	
	defer file.Close()
	
	readFile("./mygofile.txt")
}

func readFile(fileName string) {
	
	databyte, err := ioutil.ReadFile(fileName)
	
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}


func checkNilErr(err error){

	if(err != nil) {
		panic(err)
	}

}