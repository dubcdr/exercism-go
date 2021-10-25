package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	result, _ := time.Parse("1/2/2006 15:04:05", date)
	return result
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// July 25, 2019 13:45:00
	appointment, _ := time.Parse("January 2, 2006 15:04:05", date)
	return appointment.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	//Thursday, July 25, 2019 13:45:00
	appointment, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return appointment.Hour() >= 12 && appointment.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	appointment := Schedule(date)
	dayOfWeek := appointment.Format("Monday")
	month := appointment.Format("January")
	dayOfMonth := appointment.Format("2")

	return fmt.Sprintf("You have an appointment on %s, %s %s, %d, at %d:%d.", dayOfWeek, month, dayOfMonth, appointment.Year(), appointment.Hour(), appointment.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	d, _ := time.Parse("2006 01 02", fmt.Sprintf("%d 09 15", currentYear))
	return d
}
