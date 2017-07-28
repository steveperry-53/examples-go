# packages.go

Demonstrate Go packages.

## Building and running the program
    
The main source file, packages.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/13-packages
    
The packages source files are:

    $GOPATH/github.com/steveperry-53/examples-go/pkg/dog/dog.go
    $GOPATH/github.com/steveperry-53/examples-go/pkg/cat/cat.go
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/13-packages
    
The executable file, 13-packages, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/13-packages
    
Notice that a couple of static library files got built:

  $GOPATH/pkg/linux_amd64/github.com/steveperry-53/examples-go/pkg/cat.a
  $GOPATH/pkg/linux_amd64/github.com/steveperry-53/examples-go/pkg/dog.a

For a discussion of static libraries, see
[What's the usage of Golang static library file?](http://stackoverflow.com/questions/31259023/whats-the-usage-of-golang-static-library-file)

For information on the .a file extension, see
[List of archive formats](https://en.wikipedia.org/wiki/List_of_archive_formats).

The static libraries are not needed to run the executable file. The executable file is stand-alone.
