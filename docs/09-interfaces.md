# interfaces.go

Demonstrates Go interfacs.

## Building and running the program
    
The source file, variables.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/09-interfaces
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/09-interfacs
    
The executable file, 09-interfaces, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/09-interfaces
    
An *interface type* is a set of method signatures. A *value of an interface type* is
any value that implements the interface's methods.

Declare a Measure interface type:

    type Measure interface {
      Perimeter() float64
      Area() float64
    }
    
What kind of values could implement the Measure interface? Any value that could
have its area and perimeter calculated.

Declare a Rectangle type:

    type Rectangle struct {
      length float64
      width float64
    }
    
Implement the Measure interface for type Rectangle:

    func (r Rectangle) Perimeter() float64 {
      return 2*(r.length + r.width)
    }

    func (r Rectangle) Area() float64 {
      return r.length * r.width
    }
    
What do we mean when we say Rectangle implements the Measure interface?
We mean that there is a Perimeter function with a receiver of type Rectangle, and
there is an Area function with a receiver of type Rectangle.

Create a value of type Rectangle, and use the Measure interface:

    var myRect Rectangle = Rectangle{8.0, 5.0}
    
    fmt.Println(myRect.Perimeter())
    fmt.Println(myRect.Area())
    
You can also use a pointer to an interface:

    pRect := &Rectangle{length: 20.0, width: 8.0}
    
    fmt.Println((*pRect).Perimeter())
    fmt.Println(pRect.Perimeter())  // As a convenience, Go let's you skip the *.
