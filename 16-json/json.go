package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

type Thing struct {
	Name string
	Description string
}

func main() {

	// Suppose we have a JSON file that has this data.
	// Notice that the less-than sign is encoded as the escape sequence \u003c.
	stData := "{\"Name\":\"stick\",\"Description\":\"Length is \\u003c 2.\"}"
	data := []byte(stData)

	// Now suppose we unmarshal the JSON data to type Thing.
	var th Thing
	json.Unmarshal(data, &th)

	// Now th.Description has a regular less-than character. 0x3c
	// That is, the less-than character is not escaped.

	// But suppose wwe want th.Description to have an escape sequence for the less-than. \u003c

	var marshDesc []byte
	marshDesc, _ = json.Marshal(th.Description)
	strMarshDesc := string(marshDesc)

	// This is good. strMarshDesc has an escape sequence for the less-than.
	// The only problem is that the first and last characters of strMarshDesc are quote characters.

	// Trim the leading and trailing quotes.
	strMarshDesc = strings.TrimPrefix(strMarshDesc, "\"")
	strMarshDesc = strings.TrimSuffix(strMarshDesc, "\"")
	th.Description = strMarshDesc

	// Now th.Description is what we want. It is a string that has an escape sequence for the less-than.
	fmt.Println(th.Description)
}
