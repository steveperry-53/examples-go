package main

import (
  "fmt"
)

type Customer struct {
  name string
  age  int
  payments []int
}

// From the Golang Tour:
// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.

// TotalPayment is a method defined on type Customer.

// The TotalPayment method has a receiver of type Customer named c.
func (c Customer) TotalPayment() int {
  var s int = 0
  for _, p := range c.payments {
    s += p
  }
  return s
}

// The AveragePayment method has a receiver of type Customer named c.
func (c Customer) AveragePayment() float64 {
  return float64(c.TotalPayment())/float64(len(c.payments))
}

type MyString string

// The Trim method has a receiver of type MyString named s.
func (s MyString) Trim () MyString {
  l := len(s)
  return s[1:l-1]
}

type MyInt int

// The Double method has a receiver of type pointer to MyInt named i.
func (i *MyInt) Double() {
  (*i) = (*i)*2
}


func main() {
  fmt.Println("Hello methods.")

  var c1 = Customer{"Steve", 63, []int{600, 700, 900} }

  c2 := Customer{}
  c2.name = "Bob"
  c2.age = 40
  c2.payments = []int{200, 300, 400, 500, 250}

  fmt.Println(c1)
  fmt.Println(c2)

  fmt.Println(c1.TotalPayment())
  fmt.Println(c1.AveragePayment())

  var name MyString = "Steve"
  fmt.Println(name.Trim())
  fmt.Println(  (MyString("Perry")).Trim()  )

  var num MyInt = 3
  num.Double()
  fmt.Println(num)
}
