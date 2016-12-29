package main

import (
	"fmt"
)

func main() {
	// Subslices - assigment done incorrectly using simple for cycle without subslices
	// when I found correct version how to do it - there was no reason to copy-paste it
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109,
		113, 127, 131, 137, 139, 149, 151, 157,
		163, 167, 173, 179, 181, 191, 193, 197, 199}
	fmt.Println(primes)

	// Write a program to print only the primes less than 10
	// loop through the slice of primes and test if the value
	// is less than 10. When you find a value that is 10 or more
	// slice the list of primes at that point and print it.

	// Bonus: write a print only the two digit primes.

	// see src/slices/slices9/main.go for the answer

	primesLength := len(primes)

	for c := 0; c < primesLength; c++ {
		d := primes[c]
		if d < 10 {
			fmt.Printf("%d ", d)
		} else {
			fmt.Println("\nAll single-decimal primes printed.")
			break
		}
	}

	for c := 0; c < primesLength; c++ {
		d := primes[c]
		if d < 10 {
			continue
		} else if d >= 10 && d <= 99 {
			fmt.Printf("%d ", d)
		} else {
			fmt.Println("\nAll double-decimal primes printed.")
			break
		}
	}

}