package main

import (
  "fmt"
  "math"
  "strconv"
)

// Define a type. One field is a string, and one field is a function.
// Note: The Run field is a function, not a pointer to a function.

type Thing struct {
	s string
	Run func(string, int) string
}


func main() {

	// Call a function from the Go Standard Library.
	fmt.Println(math.Sqrt(37))

	// Call a function defined in this file. This function returns two int values.
	fmt.Println( AddAndMultiply(6, 7) )

	// Suppose we don't want both values returned by AddAndMultiply.
	// We only want the product, not the sum.

	_, product := AddAndMultiply(40, 20)
	fmt.Println(product) 

	// Call a function defined in this file. Returns two values: a string and an int.
	fmt.Println( swap(55, "Steve") )

	// Create an instance of type Thing.
	// Initialize with a string and a function.
	// Notice we provide the entire function body.
	t1 := Thing{
		s: "Steve",
		Run: func(name string, age int) string {
			return name + " Perry is " + strconv.Itoa(age) + "."},
	}

	// Print t1's s field and the output of t1's Run function.
	fmt.Println( t1.s, t1.Run("Edward", 64) )

	// See the function ComputeWith34, defined below.
	// What can we pass to the ComputeWith34 function?
	// We can pass any function that does something with two floats.
	// For example, return the average of the two floats.
	// Or return the area of a rectangle that has the two floats as its length and width.

	// Here's a function that returns the average of two floats.
	// We assign the function to the local variable av.
	av := func(u, v float64) float64 {
          return (u + v)/2
        }

	// Pass the av function to the ComputeWith34 function.
	fmt.Println(ComputeWith34(av))

	// Pass several other functions to the ComputeWith34 function.
	fmt.Println( ComputeWith34(Area) )
	fmt.Println( ComputeWith34(Perimeter) )
	fmt.Println( ComputeWith34(math.Pow) )
 	fmt.Println( ComputeWith34(hypot) )

	// Call ComputeWith34. This time, pass the entire body of a function.
	result := ComputeWith34(func (u float64, v float64) float64 {
		d4 := u*u*u*u + v*v*v*v
		d := math.Pow(d4, 0.25)
		return d
	})

	fmt.Println(result)
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
