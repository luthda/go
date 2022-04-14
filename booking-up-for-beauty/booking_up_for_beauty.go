package booking

import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t
}

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05" 
	t, _ := time.Parse(layout, date)

	return time.Now().After(t)
}

func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t.Hour() < 18 && t.Hour() >= 12
}

func Description(date string) string {
	layout := "January 2, 2006, at 15:04"
	dateTime := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s, %s.", 
                       dateTime.Weekday(),
                       dateTime.Format(layout))
}

func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, "2022-09-15 00:00:00")

	return t
}
