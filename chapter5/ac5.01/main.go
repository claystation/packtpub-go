package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type employee struct {
	id        int
	firstName string
	lastName  string
}

type developer struct {
	individual employee
	hourlyRate int
	workWeek   [7]int
}

func logHours(developer *developer, weekday Weekday, hours int) {
	developer.workWeek[weekday] = hours
}

func hoursWorked(developer *developer) int {
	total := 0
	for _, v := range developer.workWeek {
		total += v
	}
	return total
}

func main() {
	dev := developer{employee{1, "Knoert", "Knabbelbaard"}, 15, [7]int{}}
	logHours(&dev, Monday, 8)
	logHours(&dev, Tuesday, 8)
	fmt.Println("Hours worked on Monday: ", dev.workWeek[Monday])
	fmt.Println("Hours worked on Tuesday: ", dev.workWeek[Tuesday])
	fmt.Println("Hours worked this week: ", hoursWorked(&dev))
}
