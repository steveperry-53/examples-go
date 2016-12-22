# methods.go

Demonstrates how to define methods on types.

## Building and running the program
    
The source file, slices.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/08-methods
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/08-methods
    
The executable file, 08-methods, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/08-methods

Go does not have classes. But you can define methods on types. A method is a
function that has a receiver.

Declare a type named Customer. Each customer has a name, an age, and a slice of
ints that holds payments made by the customer:

    type Customer struct {
      name string
      age  int
      payments []int
    }

Create a TotalPayment method defined on type Customer. In other words, create
a TotalPayment function that has a receiver of type Customer.

    func (c Customer) TotalPayment() int {
      var s int = 0
      for _, p := range c.payments {
        s += p
      }
      return s
    }

Define another method on type Customer. The AveragePayment funtion has a receiver
of type Customer. A receiver has a type and a name. In this case the receiver's
name is c.

    func (c Customer) AveragePayment() float64 {
      return float64(c.TotalPayment())/float64(len(c.payments))
    }
    
 Create a value of type Customer, and call its TotalPayment method:
 
     var c1 = Customer{"Steve", 63, []int{600, 700, 900} }
     
     cmt.Println(c1.TotalPayment())
     
You can also define methods on types that are not structs. But the type
must be defined in your package. For example, if you define a type MyString, you
can define methods on it.

    type MyString string

    func (s MyString) Trim () MyString {
      l := len(s)
      return s[1:l-1]
    }
    
    fmt.Println( (MyString("Perry")).Trim() )
     
 

