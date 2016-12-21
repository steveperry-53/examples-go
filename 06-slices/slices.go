package main

import "fmt"

var names [6]string
var ages [6]int

func main() {
  fmt.Println("Hello slices")

  names[0] = "Steve"
  names[1] = "Bob"
  names[2] = "Sam"
  names[3] = "Ted"
  ages[0] = 60
  ages[1] = 25
  ages[2] = 40
  ages[3] = 32
  
  fmt.Println(names[1], ages[1])
  fmt.Println(names)
  fmt.Println(ages)

  weights := [5]int{170, 150, 120, 210}

  fmt.Println(weights)

  // Create some slices.
  var sn []string = names[1:3] // Half-open interval, 1 <= index < 3. https://blog.golang.org/go-slices-usage-and-internals
  var sa []int = ages[1:3] 
  var sw []int = weights[1:3]

  fmt.Println(sn)
  fmt.Println(sa)
  fmt.Println(sw)

  // Slice literals
  var pets = []string{"Fluff", "Spot", "Rover"}
  trees := []string{"Pine", "Fir", "Redwood", "Spruce"}

  fmt.Println(pets)
  fmt.Println(trees)

  type Dog struct {
    name string
    age int
  }

  dogs := []Dog{ {"Rover", 5}, {"Growler", 15} }

  fmt.Println(dogs)

  // Shortcuts

  fmt.Println(pets[:2])
  fmt.Println(trees[2:])

  fmt.Println(names[:3])

  // You can use "for ... range" to iterate over a slice.
  // range returns two values: the index and the element.
 
  for i, v := range pets {
    fmt.Println(v, "ate", i, "rats.")
  }

  for i := range pets {
    fmt.Println(i)
  }

  for _, v := range pets {
    fmt.Println(v)
  }

}
