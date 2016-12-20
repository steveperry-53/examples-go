# structs.go

Demonstrates how to use structs.

Imports the
[fmt](https://golang.org/pkg/fmt/)
package from the
[Go Standard Library](https://golang.org/pkg/#stdlib).
Calls the
[fmt.Println](https://golang.org/pkg/fmt/#Println) function.

## Building and running the program
    
The source file, variables.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/05-structs
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/05-structs
    
The executable file, 02-variables, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/05-structs

A *struct* is a colletion of *fields*.

The program declares a type named Rectangle:

    type Rectangle struct {
      width int
      length int
    }

There are several ways to initialize a struct.

    r := Rectangle{3, 5}
    var s = Rectangle{7, 10}
    var t = Rectangle{width: 12, length: 15}
    var u = Rectangle{width: 20}
    var v = &Rectangle{length: 30}
    w := &Rectangle{width: 40}

Notice that v is the address of a Rectangle.
You can refer to fields of v two ways:

    v.length
    (*v).length

The program also declares a type named Command. This is is similar
to the cobra.Command type in
[create_secret.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/create_secret.go).

A Command struct has several string fields and one field of type
func(code string). You can call the function like this:

    cmd := &Command{ ... }
    cmd.Run("message1")
    (*cmd).Run("message2")



