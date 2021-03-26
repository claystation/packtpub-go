package main

import "fmt"

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value:\t", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point:\t", *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:\t", count)

	add5Point(&count)
	fmt.Println("Add5Point post:\t", count)
}
