package main

import (
	"errors"
	"fmt"
)

func goodCreditScore(score int) bool {
	return score >= 450
}

func interestRate(goodCredit bool) float64 {
	if goodCredit {
		return 15
	} else {
		return 20
	}
}

func monthlyPaymentAllowed(income float64, monthlyPayment float64, goodCreditScore bool) bool {
	if goodCreditScore {
		return monthlyPayment <= (income * 0.20)
	} else {
		return monthlyPayment <= (income * 0.15)
	}
}

func validateLoan(creditScore int, income float64, amount float64, terms float64) error {
	if creditScore <= 0 || income <= 0 || amount <= 0 || terms <= 0 {
		errors.New("invalid parameters")
	}
	if int(terms)%12 != 0 {
		errors.New("terms not divisible by 12")
	}

	goodCredit := goodCreditScore(creditScore)
	interest := interestRate(goodCredit)
	interestAmount := (amount * (interest / 100)) / terms
	monthlyPayment := (amount / terms) + interestAmount

	fmt.Println("Credit Score : ", creditScore)
	fmt.Println("Income : ", income)
	fmt.Println("Loan Amount : ", amount)
	fmt.Println("Loan Term : ", terms)
	fmt.Println("Monthly Payment : ", monthlyPayment)
	fmt.Println("Rate : ", interest)
	fmt.Println("Total Cost : ", amount*(interest/100))
	fmt.Println("Approved : ", monthlyPaymentAllowed(income, monthlyPayment, goodCredit))

	return nil
}

func main() {

	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	validateLoan(500, 1000, 1000, 24)

	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	validateLoan(350, 1000, 10000, 12)
}
