package main

import (
	"fmt"
	"strconv"
)

func main() {

	firstName := "John"
	lastName := "Doe"
	age := 31
	allergic := true

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(strconv.Itoa(age))
	fmt.Println(allergic)

}
