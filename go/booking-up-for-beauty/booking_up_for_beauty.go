package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string, format ...string) time.Time {
	var _fmt string
	if len(format) > 0 {
		_fmt = format[0]
	} else {
		_fmt = "1/02/2006 15:04:05"
	}

	t, _ := time.Parse(_fmt, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return time.Now().After(Schedule(date, "January 2, 2006 15:04:05"))
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	hour := Schedule(date, "Monday, January 2, 2006 15:04:05").Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	date = Schedule(date, "1/2/2006 15:04:05").Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", date)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
