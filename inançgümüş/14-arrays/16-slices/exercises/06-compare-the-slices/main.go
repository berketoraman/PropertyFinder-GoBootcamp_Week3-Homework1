// EXERCISE: Compare the slices
//  1. Split the namesA string and get a slice
//  2. Sort all the slices
//  3. Compare whether the slices are equal or not

// EXPECTED OUTPUT
//   They are equal.
// HINTS
//   1. strings.Split function splits a string and
//      returns a string slice
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesC := strings.Split(namesA, ", ")

	sort.Strings(namesC)
	sort.Strings(namesB)

	if len(namesC) == len(namesB) {
		for i := range namesC {
			if namesC[i] != namesB[i] {
				return
			}
		}
		fmt.Println("They are equal.")
	}
}
