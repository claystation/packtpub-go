package main

import (
	"fmt"

	"github.com/claystation/packtpub-go/chapter8/ac8.01/payroll"
)

var d payroll.Developer
var m payroll.Manager

func init() {
	fmt.Println("Welcome to Payroll Solutions")
}

func init() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Initializing variables")

	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d = payroll.Developer{Individual: payroll.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m = payroll.Manager{Individual: payroll.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
}

func main() {
	d.ReviewRating()
	printPayment(d)
	printPayment(m)
}

func printPayment(p payroll.Payer) {
	name, salary := p.Pay()
	fmt.Printf("%s got paid %.2f for the year.\n", name, salary)
}
