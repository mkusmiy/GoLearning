package main

import (
	"fmt"
)

func main() {
	//Slice exercises
	// Declare a variable called e which is a slice of 5 int.
	e := make([]int, 5)
	// Declare a variable called f which is a slice of 9 float64.
	f := make([]float64, 9)
	// Declare a variable called s which is a slice of 4 string.
	s := make([]string, 4)

	fmt.Println(len(e), len(f), len(s))
}