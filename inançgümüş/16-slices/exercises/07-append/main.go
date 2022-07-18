// EXERCISE: Append
//  Please follow the instructions within the code below.

// EXPECTED OUTPUT
//  They are equal.

// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal

package main

import (
	"bytes"
	"fmt"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	if bytes.Equal(png, header) {
		fmt.Println("They are equal")
	}
}
