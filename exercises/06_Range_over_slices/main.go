package main

import (
	"fmt"
)

func main() {
	//Range over slices
	primes2 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	// This program has two problems.
	// The first problem is the output reports the 1st prime
	// as the 0 th prime.
	// The second problem is the suffix the 1st prime should be
	// 1 st, not 1 th.
	// Can you fix both of these problems.

	for i, p := range primes2 {
		switch i {
		case 0:
			fmt.Println("The", i+1, "st prime is", p)
		case 1:
			fmt.Println("The", i+1, "nd prime is", p)
		case 2:
			fmt.Println("The", i+1, "rd prime is", p)
		default:
			fmt.Println("The", i+1, "th prime is", p)
		}

	}
}