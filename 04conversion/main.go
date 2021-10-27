package main

import(
  "fmp"
  "bufio"
  "os"
)

func main() {

  fmp.Println("Welcome to our pizza app")
  fmp.Println("Please rate our pizza between 1 and 5")

  reader := bufio.NewReader(os.Stdin)

  input, _ := reader.ReadString('\n')
  fmt.Println("Thanks for review", input)

  strconv.Parse

}