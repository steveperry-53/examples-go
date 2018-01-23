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

	// To use the description in HTML, we need to replace the unicode excape sequence with an HTML entities.
	strMarshDesc = strings.Replace(strMarshDesc, "\\u003c", "&lt;", -1)
	th.Description = strMarshDesc

	// Now th.Description is what we want. It is a string that has an HTML entity for the less-than.
	fmt.Println(th.Description)


	// Next, suppose we have a JSON file that has this data.
	stData2 := "{\"Name\":\"Endpoints\", \"Description\":\"Endpoints is a collection. Example:\\n  Name: \\\"mysvc\\\"\"}"
	data2 := []byte(stData2)

	// Unmarshal the JSON data to type Thing.
	var th2 Thing
	json.Unmarshal(data2, &th2)

	// Now th2.Description has  regular quote characters and a regular newline character.
	// That is, the quotes are not escaped. The newline character is not escaped.

	// We could use th2.Description in an HTML document. We wouldn't get new lines at the newline
	// characters, but other than that, the output would be OK. We wouldn't see weird characters
	// in the rendered HTML. We wouldn't see \" or \n in the rendered HTML.

	// But we're going to call json.Marshal on the description because there might be angle brackets
	// in the description. And we need to convert the angle brackets back to Unicode escape sequences
	// and then to HTML entities.

	// This messes up our quotes and newlines. Now we're going to get \n and \" again in the rendered HTML.

	// We could use strings.Replace to replace \\n with \n and \\\" with \".
	// But it seems like we're goint in circles.

	// Is there some other way to deal with the angle brackets in descriptios?
	// Could I skip the json.Marshal step, and instead just replace < with &lt; and replace > with &gt;?
}

