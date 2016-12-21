# slices.go

Demonstrates how to use arrays and slices.

Imports the
[fmt](https://golang.org/pkg/fmt/)
package from the
[Go Standard Library](https://golang.org/pkg/#stdlib).
Calls the
[fmt.Println](https://golang.org/pkg/fmt/#Println) function.

## Building and running the program
    
The source file, slices.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/06-slices
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/06-slices
    
The executable file, 06-slices, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/06-slices

Arrays have fixed length. Slices can be sized dynamically.

Declare an array of 6 strings and an array of 6 ints:

    var names [6]string
    var ages [6]int
  
Use [] to assign values to elements of the arrays:

    names[0] = "Steve"
    names[1] = "Bob"
    ...
    ages[0] = 60
    ages[1] = 25
    ...

Use [] to access elements of the arrays:

    fmt.Println(names[1], ages[1])

Use a short variable declaration to declare and initialize
and array of 5 integers:

    weights := [5]int{170, 150, 120, 210}

Note: The last element of the array is initialized with 0.

Create a slice from an existing array:

    var sn []string = names[1:3]

Slices are based on a half-open interval. In the preceding example,
1 <= index < 3. So the slice has elements 1 and 2 of the names array.

Create a new slice of strings, and initialize it with a slice literal:

    var pets = []string{"Fluff", "Spot", "Rover"}

Declare a type Dog, and create a slice of Dogs:

    type Dog struct {
    name string
    age int
    }

    dogs := []Dog{ {"Rover", 5}, {"Growler", 15} }

You can use for ... range to iterate over a slice:

    for i, v := range pets {
      fmt.Println(v, "ate", i, "rats.")
    }




     
