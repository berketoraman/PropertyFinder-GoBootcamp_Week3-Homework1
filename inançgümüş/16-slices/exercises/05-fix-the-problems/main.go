// EXERCISE: Fix the problems
//  1. Uncomment the code
//  2. Fix the problems
//  3. BONUS: Simplify the code

// EXPECTED OUTPUT
//   "Einstein and Shepard and Tesla"
//   ["Fire" "Kafka's Revenge" "Stay Golden"]
//   [1 2 3 5 6 7 8 9]

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	names := []string{"Einstein", "Shepard", "Tesla"}

	books := []string{"Stay Golden", "Fire", "Kafka's Revenge"}
	sort.Strings(books)

	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}
	sort.Ints(nums[:])

	fmt.Printf("%q\n", strings.Join(names, " and "))
	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
