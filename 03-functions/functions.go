package main

import (
  "fmt"
  "math"
  "strconv"
)

type Thing struct {
  s string
  Run func(string, int) string
}


func main() {
  fmt.Println(math.Sqrt(37))
  fmt.Println( AddAndMultiply(6, 7) )
  fmt.Println( swap(55, "Steve") )

  av := func(u, v float64) float64 {
          return (u + v)/2
        }

  fmt.Println( ComputeWith34(Area) )
  fmt.Println( ComputeWith34(Perimeter) )
  fmt.Println( ComputeWith34(math.Pow) )
  fmt.Println( ComputeWith34(hypot) )
  fmt.Println( ComputeWith34(av) )

  t1 := Thing{
    s: "Steve",
    Run: func(name string, age int) string {
      return name + " Perry is " + strconv.Itoa(age) + "."
    },
  }

  fmt.Println( t1.s, t1.Run("Edward", 64) )

}


func AddAndMultiply(x, y int) (int, int) {
  return x + y, x * y
}


func swap(age int, name string) (string, int) {
  return name, age
}


func Area(x, y float64) float64 {
  return x*y
}


func Perimeter(x, y float64) float64 {
  return 2*(x + y)
}


func ComputeWith34(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

var hypot = func(a, b float64) float64 {
              return math.Sqrt(a*a + b*b)
            }
