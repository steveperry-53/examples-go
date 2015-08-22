package main

import "fmt"

func main() {

  var a, b, c, d int
  var fir, lst string

  a = 3
  b = 2
  c, d = a + b, a * b

  fmt.Println(c, d)

  fir, lst = "Steve", "Perry"

  fmt.Println(fir, lst)

  var e = 8
  f := 9

  var g, h = 10, 11

  fmt.Println(e, f, g, h)

  var i bool = true
  var j bool = false

  fmt.Println(i, j)

  var k float32 = 3.14
  var l float64 = 2.71

  var m complex64 = 5 + 3i

  n := 6 + 4i

  fmt.Println(k, l, m, n)
}

//////////////////////////////////////////

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// float32 float64

// complex64 complex128
