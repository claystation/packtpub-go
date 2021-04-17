package main

import "fmt"

func main () {
	daysOfWeek := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	sundayStart := append(daysOfWeek[6:], daysOfWeek[:6]...)

	fmt.Println(sundayStart)
}