package main

import (
	"fmt"
	"github.com/steveperry-53/examples-go/13-packages/pkg/dog"
)

func main() {
	myDog := Dog{"fido", 6, 40}
	fmt.Println(myDog.Weight)
}
