package main

import (
  "fmt"
  "math"
)

// From the Golang Tour
// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

// What does it mean for a type to implement those methods?
// I think it means that those methods exist with the type as receiver.



type Measure interface {
  Perimeter() float64
  Area() float64
}


type Rectangle struct {
  length float64
  width float64
}

type Circle struct {
  radius float64
}

type Triangle struct {
  base float64
  height float64
}


// Rectangle implementation of the Measure interface

func (r Rectangle) Perimeter() float64 {
  return 2*(r.length + r.width)
}

func (r Rectangle) Area() float64 {
  return r.length * r.width
}


// Circle implementation of the Measure interface

func (c Circle) Perimeter() float64 {
  return 2 * c.radius * math.Pi
}

func (c Circle) Area() float64 {
  return math.Pi * c.radius * c.radius
}


// Triangle implementation of the Measure interface

func (t Triangle) Perimeter() float64 {
  return 0.0
}

func (t Triangle) Area() float64 {
  return (t.base * t.height)/2.0
}


func main() {
  fmt.Println("Hello interfaces.")


  var myRect Rectangle = Rectangle{8.0, 5.0}
  fmt.Println(myRect.Perimeter())
  fmt.Println(myRect.Area())

  myCirc := Circle{5.0}
  fmt.Println(myCirc.Perimeter())
  fmt.Println(myCirc.Area())

  myTri := Triangle{base: 10.0, height: 4.0}
  fmt.Println(myTri.Perimeter())
  fmt.Println(myTri.Area())

  pRect := &Rectangle{length: 20.0, width: 8.0}
  fmt.Println((*pRect).Perimeter())
  fmt.Println((*pRect).Area())
  fmt.Println(pRect.Perimeter())  // As a convenience, Go let's us skip the *.
  fmt.Println(pRect.Area())

  // What is an interface value? Maybe this:
  var myIface Measure = Rectangle{80.0, 90.0}

  fmt.Println(myIface)
  val, ok := myIface.(Rectangle)
  fmt.Println(val, ok)

  val2, ok2 := myIface.(Circle)
  fmt.Println(val2, ok2)

  fmt.Printf("%v %T\n", myIface, myIface)

  fmt.Printf("%v %T\n", myRect, myRect)
}
