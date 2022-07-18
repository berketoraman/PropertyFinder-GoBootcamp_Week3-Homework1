/*Sieve of Eratosthenes is an algorithm for generating all prime numbers less 
than or equal to a given number N
*/
package main

import (
		"fmt"
)

// How we display the output.
func display(num int, prime []bool) {
		fmt.Printf("Primes less than %d : ", num) //This line shows how the output will be when we call the sieve function
		for i := 2; i <= num; i++ { //The smallest prime number is 2.
				if prime[i] == false {
						fmt.Printf("%d ", i)
				}
		}
		fmt.Println()
}

func sieve(num int) {
    prime := make([]bool, num+1)

	// initialize everything with false first(not crossed)
	for i := 0; i < num+1; i++ {
				prime[i] = false
	}

	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true // cross
			}
		}

	}
    display(num, prime)
}

//Run the examples in the main function to see the results
func main() {
		sieve(10)
		sieve(30)
		sieve(50)
		sieve(80)
}