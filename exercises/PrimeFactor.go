//"Prime Factorization" is finding which prime numbers multiply together to make the original number
package main

import (
	"fmt"
)

// Get all prime factors of a given number n
func PrimeFactors(n int) (pfactors []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfactors = append(pfactors, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfactors = append(pfactors, i)
			n = n / i
		}
	}

/* This condition is to handle the case when n is a prime number
greater than 2
*/	
    if n > 2 {
		pfactors = append(pfactors, n)
	}

	return
}

func main() {
	fmt.Println(PrimeFactors(17))
	fmt.Println(PrimeFactors(24))
	fmt.Println(PrimeFactors(480))
	fmt.Println(PrimeFactors(79))
}
