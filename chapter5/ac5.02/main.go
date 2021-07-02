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

type Employee struct {
	id        int
	firstName string
	lastName  string
}

type Developer struct {
	individual Employee
	hourlyRate int
	workWeek   [7]int
}

func logHours(developer *Developer, weekday Weekday, hours int) {
	developer.workWeek[weekday] = hours
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(hours int) int {
		total += hours
		return total
	}
}

func hoursWorked(developer *Developer) int {
	total := 0
	for _, v := range developer.workWeek {
		total += v
	}
	return total
}

func (d *Developer) PayDay() (int, bool) {
	if hoursWorked(d) > 40 {
		hoursOver := hoursWorked(d) - 40
		overTime := hoursOver * 2 * d.hourlyRate
		regularPay := 40 * d.hourlyRate
		return regularPay + overTime, true
	}
	return hoursWorked(d) * d.hourlyRate, false
}

func main() {
	dev := Developer{Employee{1, "Knoert", "Knabbelbaard"}, 15, [7]int{}}
	nonLogged := nonLoggedHours()
	fmt.Println("Tracking hours worked so far", nonLogged(2))
	fmt.Println("Tracking hours worked so far", nonLogged(3))
	fmt.Println("Tracking hours worked so far", nonLogged(5))
	logHours(&dev, Monday, 8)
	logHours(&dev, Tuesday, 10)
	logHours(&dev, Wednesday, 10)
	logHours(&dev, Thursday, 10)
	logHours(&dev, Friday, 6)
	logHours(&dev, Saturday, 8)
	logHours(&dev, Sunday, 0)
	dev.PayDetails()
}

func (d *Developer) PayDetails() {
	for i, v := range d.workWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}
	fmt.Printf("\nHours worked this week:  %d\n", hoursWorked(d))
	pay, overtime := d.PayDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
	fmt.Println()
}
