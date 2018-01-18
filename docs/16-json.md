# json.go

Demonstrates json.Unmarshal.

## Building and running the program
    
The source file, json.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/16-json
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/16-json
    
The executable file, 16-json, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/16-json

## Discussion

This program does the following:

1. Create a byte slice that has an escaped less than character. \uc003
1. Unmarshal the byte slice into an value of type Thing.
1. Display each byte of the original byte array and each byte of the Thing's Description field.
1. Notice that the Thing's description field has an actual less than character. It is no longer escaped.
