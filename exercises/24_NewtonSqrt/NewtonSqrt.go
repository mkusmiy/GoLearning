package main

import (
	"fmt"
	"math"
)

func NSqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = 0.5 * (z + x/z)
	}
	return z
}

func main() {
	fmt.Println(NSqrt(4))
	fmt.Println(math.Sqrt(4))
}
