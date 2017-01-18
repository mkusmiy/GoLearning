package main

import "fmt"

// exercise to use anonymous and recursive
// function (closure)

type recfunc func(recfunc, int) int

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(recfunc, int) int {
	fib := 0
	return func(f recfunc, x int) int {
		if x == 0 {
			return 0
		} else if x == 1 {
			return 1
		} else {
			return f(f, x-1) + f(f, x-2)
		}
		return fib
	}
}

// easier way to do it (from internet)
func fibonacci1() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Printf("%v ", f(f, i))
	}
	fmt.Println("")
}
