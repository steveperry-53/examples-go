Prints Hello World. Imports the
[fmt](https://golang.org/pkg/fmt/)
package from the
[Go Standard Library](https://golang.org/pkg/#stdlib).
Calls the [fmt.Println](https://golang.org/pkg/fmt/#Println) function.

### Building and running the program

On your development machine, create a $HOME/work directory.
Set $GOPATH to $HOME/work.

Navigate to $GOPATH.

    mkdir src
    cd src
    go get github.com/steveperry-53/examples-go
    
The source file, hello.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/01-hello-world
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/01-hello-world
    
The executable file, 01-hello-world, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/01-hello-world
    
### Other go commands

* go build
* go run 
    
