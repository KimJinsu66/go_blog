// https://go-tour-ko.appspot.com/basics/5
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func minus(x, y int) int {
  return x - y
}

func main() {
  fmt.Println(add(42, 13))
  fmt.Println(minus(42, 13))
}
