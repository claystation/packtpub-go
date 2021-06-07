package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)


type directDeposit struct {
	lastName string
	firstName string
	bankName string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}

func (d *directDeposit) validateLastName() error {
	if len(d.lastName) == 0 {
		return ErrInvalidLastName
	}
	return nil
}


func main() {

	d := directDeposit{
		lastName: "",
		firstName: "Abe",
		bankName: "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	if err := d.validateLastName(); err != nil {
		fmt.Println(err)
	}

	if err := d.validateRoutingNumber(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("*************************************")
	fmt.Println("Last name:", d.lastName)
	fmt.Println("First name:", d.firstName)
	fmt.Println("Bank Name:", d.bankName)
	fmt.Println("Routing Number:", d.routingNumber)
	fmt.Println("Account Number:", d.accountNumber)

}
