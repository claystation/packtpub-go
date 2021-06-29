package main

import (
	"fmt"
)

type employee struct {
	id int
	firstName string
	lastName string
}

type developer struct {
	individual employee
	hourlyRate float64
	hoursWorkedInYear float64
	review map[string]interface{}
}

type manager struct {
	individual employee
	salary float64
	commissionRate float64
}

type payer interface {
	pay() (string, float64)
}

func (e employee) fullName() string {
	return fmt.Sprintf("%s %s", e.firstName, e.lastName)
}

func (d developer) pay() (string, float64) {
	return d.individual.fullName(), d.hourlyRate * d.hoursWorkedInYear
}

func (m manager) pay() (string, float64) {
	return m.individual.fullName(), m.salary + (m.salary * m.commissionRate)
}

func payDetails(p payer) {
	name, pay := p.pay()
	fmt.Printf("%s got paid %.2f for the year. \n", name, pay)
}

func review(d developer) {
	rating := float64(0)
	for _, v := range d.review {
		switch v.(type) {
		case int:
			rating += float64(v.(int))
		case string:
			switch v {
			case "Excellent":
				rating += 5
			case "Good":
				rating += 4
			case "Fair":
				rating += 3
			case "Poor":
				rating += 2
			case "Unsatisfactory":
				rating += 1
			}
		}
	}
	fmt.Printf("%s got a review rating of %.2f \n", d.individual.fullName(), rating / float64(len(d.review)))
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := developer{individual: employee{id: 1, firstName: "Eric", lastName: "Davis"}, hourlyRate: 35, hoursWorkedInYear: 2400, review: employeeReview}
	m := manager{individual: employee{id: 2, firstName: "Mr.", lastName: "Boss"}, salary: 150000, commissionRate: .07}
	payDetails(d)
	review(d)
	payDetails(m)
}