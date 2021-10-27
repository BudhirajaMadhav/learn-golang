package main

import (
  "fmt"
  "bufio"
  "os"
  )

func main() {

  welcome := "welcome to user input"
  fmt.Println(welcome)
  
  reader := bufio.NewReader(os.Stdin)


  // treat errors like variables or true false

  // comma ok || comma err syntax
  fmt.Println("Please give review")
  // we don't care about the returned error here so put "_". else err.
  input, _ := reader.ReadString('\n')
  fmt.Println("thank you for your rating of", input)
  fmt.Printf("Type of this rating is %T mf", input)
}