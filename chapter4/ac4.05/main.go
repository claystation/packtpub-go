package main

import "fmt"

func main() {
	goodBad :=[]string{
		"Good",
		"Good",
		"Bad",
		"Good",
		"Good",
	}

	fmt.Println(append(goodBad[:2], goodBad[3:]...))
}