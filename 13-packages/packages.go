package main

import (
	"fmt"
	"github.com/steveperry-53/examples-go/pkg/dog"
	"github.com/steveperry-53/examples-go/pkg/cat"
)

func main() {
	myDog := dog.Dog{"Fido", 50}
	fmt.Println(myDog)
	fmt.Println(myDog.Weight())

	myCat := cat.Cat{"Fluff", 7}
	fmt.Println(myCat)
	fmt.Println(myCat.Age())
}
