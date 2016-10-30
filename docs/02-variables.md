# variables.go

Demonstrates how to declare, initialize, and print variables
of several different types.

Imports the
[fmt](https://golang.org/pkg/fmt/)
package from the
[Go Standard Library](https://golang.org/pkg/#stdlib).
Calls the
[fmt.Println](https://golang.org/pkg/fmt/#Println) function.

## Building and running the program
    
The source file, variables.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/02-variables
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/02-variables
    
The executable file, 02-variables, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/01-variables

Go has these
[basic types](http://tour.golang.org/basics/11):

* bool
* string
* int  int8  int16  int32  int64
* uint uint8 uint16 uint32 uint64 uintptr
* byte // alias for uint8
* rune // alias for int32
* float32 float64
* complex64 complex128

To declare a variable, use the `var` keyword.
The type goes after the variable name:

    var x int
    var y int = 5

You can declare and initialize several variables on one line:

    var a, b int

Use = for regular variable assignment:

    a = 7

You can assign values to several variables on one line:

    b, x = 8, 9

Use := for short assignment:

    c := 10
    firstName := "Steve"
    

## See also

[Golang book, Variables](https://www.golang-book.com/books/intro/4)

