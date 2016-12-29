package main

import (
	"fmt"
)

func main() {
	//What time is it now
	now := time.Now()
	fmt.Println("The current time is", now.Hour(), ":", now.Minute())
}