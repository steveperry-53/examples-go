package main

import "fmt"

func main() {

  for i := 1; i <= 10; i++ {
    fmt.Println(i*i)
  }

  fmt.Println()
  fmt.Println()

  bal := 100.0
  fmt.Println(bal)

  for ; bal < 200 ; {
    bal = bal * 1.08
    fmt.Println(bal)
  }

  sum := 1

  for sum < 100 {
    sum = sum + sum
    fmt.Println(sum)
  }


}


