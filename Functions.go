package main

import "fmt"

func add (x int, y int) int { //to shorten do  func add (x, y int)
  return x + y
  
}

func main() {
  fmt.Println(add(42, 13))
  
}
