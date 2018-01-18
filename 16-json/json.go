package main

import (
	"fmt"
	"encoding/json"
)

type Thing struct {
	Name string
	Description string
}

func main() {
	fmt.Println("Hello json")

	var th Thing

	stData := "{\"Name\":\"stick\",\"Description\":\"Length is \\u003c 2.\"}"

	data := []byte(stData)

	json.Unmarshal(data, &th)

	fmt.Println()
	fmt.Println("stData")
	for j := 0; j < len(stData); j++ {
		fmt.Printf("%c %x\n", stData[j], stData[j])
	}

	fmt.Println()
	fmt.Println("data")
	for j:= 0; j < len(data); j++ {
		fmt.Printf("%c %x\n", data[j], data[j])
	}

	fmt.Println()
	fmt.Println("th.Description")
	for j:= 0; j < len(th.Description); j++ {
		fmt.Printf("%c %x\n", th.Description[j], th.Description[j])
	}
}
