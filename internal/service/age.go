package service

import (
	"time"
)

// CalculateAge calculates the age based on the date of birth and the current time.
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
