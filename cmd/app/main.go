package main

import (
	"fmt"
	"unicode"
)

func main() {

	var s string = "saveChangesInTheEditor"
	fmt.Println(camelcase(s))

}

func camelcase(s string) int32 {

	// declare word counter
	var i int32 = 0

	// foreach letter in the string, if its uppercase, increase counter
	for _, l := range s {
		if unicode.IsUpper(l) {
			i++
		}
	}

	// first word will always be lowercase
	return i + 1
}
