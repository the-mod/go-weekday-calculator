package main

import (
	"strings"
	"time"
)

func getNumberOfDayOfWeek(day string) int {
	day = strings.ToLower(day)
	var days [7]string
	// TODO IOTA
	days[0] = "monday"
	days[1] = "tuesday"
	days[2] = "wednesday"
	days[3] = "thursday"
	days[4] = "friday"
	days[5] = "saturday"
	days[6] = "sunday"

	for key, value := range days {
		if day == value {
			return key
		}
	}
	return -1
}

// gets the format value of corresponding pattern
// see: https://golang.org/src/time/format.go
func getDateFormat(inputFormat string) string {
	r := strings.NewReplacer(
		"dd", "02",
		"MM", "01",
		"YYYY", "2006",
		"YY", "06",
	)
	return r.Replace(inputFormat)
}

// GetFormatedDateString calculates dates based on the current calendar week and given pattern
// example: current calendarweek 51 produces GetFormatedDateString("Monday", "dd.MM.YYYY") => 14.12.2020
func GetFormatedDateString(dayOfWeekString string, pattern string) string {
	dayOfWeek := getNumberOfDayOfWeek(dayOfWeekString)
	now := time.Now()
	currentWeekdayString := now.Weekday().String()
	currentWeekday := getNumberOfDayOfWeek(currentWeekdayString)

	// if current day of week is equals or greater as given dayOfWeek then substract
	if currentWeekday >= dayOfWeek {
		diff := currentWeekday - dayOfWeek
		newDate := time.Now().AddDate(0, 0, -diff)
		return newDate.Format(getDateFormat(pattern))
	} else {
		// else add
		diff := dayOfWeek - currentWeekday
		newDate := time.Now().AddDate(0, 0, diff)
		return newDate.Format(getDateFormat(pattern))
	}
}
