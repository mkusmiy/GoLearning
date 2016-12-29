package main

import (
	"fmt"
)

func main() {
	//Short declaration syntax
	i := 0
	for i = 1; i < 11; i++ {
		fmt.Println(i)
	}

	j := 10
	for ; j > 0; j-- {
		fmt.Println(j)
	}
}