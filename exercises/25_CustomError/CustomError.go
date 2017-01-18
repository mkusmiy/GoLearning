package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	msg := fmt.Sprint("Cannot Sqrt negative number: ", float64(e))
	return msg
}

func NSqrt(x float64) (float64, error) {
	z := 1.0
	if x == 0 {
		z = 0
	} else if x > 0 {
		for i := 0; i < 10; i++ {
			z = 0.5 * (z + x/z)
		}
	} else {
		return 0, ErrNegativeSqrt(x)
	}
	return z, nil
}

func main() {
	var n float64 = -2
	fmt.Println(NSqrt(n))
	fmt.Println(math.Sqrt(n))

}
