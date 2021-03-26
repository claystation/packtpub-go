package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You": 3,
		"Give": 2,
		"Never": 1,
		"Up":  4,
	}

	var count int
	var word string
	for key, value := range words {
		if value > count {
			count = value
			word = key
		}
	}
	fmt.Printf("Most counted word: #%v with count: #%v", word, count)

}