package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(math.Sqrt(37))
  fmt.Println( AddAndMultiply(6, 7) )
  fmt.Println( swap(55, "Steve") )
}


func AddAndMultiply(x, y int) (int, int) {
  return x + y, x * y
}


func swap(age int, name string) (string, int) {
  return name, age
}
