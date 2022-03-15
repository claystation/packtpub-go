package main

import (
	"errors"
	"log"
	"strings"
	"unicode"
)

var (
	ErrInvalidSSNLength  = errors.New("Invalid SSN Length")
	ErrInvalidSSNNumbers = errors.New("Invalid SSN Numbers")
	ErrInvalidSSNPrefix  = errors.New("Invalid SSN Prefix")
	ErrInvalidDigitPlace = errors.New("Invalid SSN Digit Place")
)

func ssnLenghtChecker(ssn string) error {
	ssn = strings.TrimSpace(ssn)
	if len(ssn) != 9 {
		return ErrInvalidSSNLength
	}
	return nil
}

func ssnNumberChecker(ssn string) error {
	for _, v := range ssn {
		if !unicode.IsNumber(v) {
			return ErrInvalidSSNNumbers
		}
	}
	return nil
}

func ssnPrefixChecker(ssn string) error {
	if strings.HasPrefix(ssn, "000") {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func ssnInvalidPlace(ssn string) error {
	if string(ssn[0]) == "9" {
		if string(ssn[4]) != "7" && string(ssn[4]) != "9" {
			return ErrInvalidDigitPlace
		}
	}
	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}
	log.Printf("Validating data: %v", validateSSN)

	for k, v := range validateSSN {
		log.Printf("Validate data %s %d of %d", v, k+1, len(validateSSN))
		v = strings.Replace(v, "-", "", -1)
		err := ssnLenghtChecker(v)
		if err != nil {
			log.Print(err)
		}
		err = ssnNumberChecker(v)
		if err != nil {
			log.Print(err)
		}
		err = ssnPrefixChecker(v)
		if err != nil {
			log.Print(err)
		}
		err = ssnInvalidPlace(v)
		if err != nil {
			log.Print(err)
		}

	}
}
