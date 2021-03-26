package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(value int) string {
	if value%3 == 0 && value%5 == 0 {
		return "FizzBuzz"
	} else if value%3 == 0 {
		return "Fizz"
	} else if value%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(value)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
