package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(password string) bool {
	passwordRune := []rune(password)

	if len(passwordRune) < 8 || len(passwordRune) > 15 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, value := range passwordRune {
		if unicode.IsUpper(value) {
			hasUpper = true
		}
		if unicode.IsLower(value) {
			hasLower = true
		}
		if unicode.IsNumber(value) {
			hasNumber = true
		}
		if unicode.IsPunct(value) || unicode.IsSymbol(value) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("Password good")
	} else {
		fmt.Println("Password bad")
	}

	if passwordChecker("This!I5A") {
		fmt.Println("Password good")
	} else {
		fmt.Println("Password bad")
	}
}
