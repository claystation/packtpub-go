package payroll

import (
	"fmt"

	"github.com/claystation/packtpub-go/chapter8/ac8.01/review"
)

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	fullName := d.Individual.FullName()
	return fullName, d.HourlyRate * d.HoursWorkedInYear
}

func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := review.OverallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.Individual.FullName(), averageRating)
	return nil
}
