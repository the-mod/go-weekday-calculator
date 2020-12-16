package main

import (
	"testing"
	"time"
)

func TestGetFormatedDateStringOfCurrentTime(t *testing.T) {
	currentWeekdayString := time.Now().Weekday().String()

	expected := time.Now().Format("02.01.2006")

	result := GetFormatedDateString(currentWeekdayString, "dd.MM.YYYY")
	if result != expected {
		t.Errorf("Expected: \"%v\", Current: \"%v\"", expected, result)
	}
}

func TestGetFormatedDateStringOfCurrentTimeWithISO8601FormatYYYYMMdd(t *testing.T) {
	currentWeekdayString := time.Now().Weekday().String()
	expected := time.Now().Format("2006-01-02")

	result := GetFormatedDateString(currentWeekdayString, "YYYY-MM-dd")
	if result != expected {
		t.Errorf("Expected: \"%v\", Current: \"%v\"", expected, result)
	}
}

func TestGetFormatedDateStringOfCurrentTimeWithISO8601FormatMMddYY(t *testing.T) {
	currentWeekdayString := time.Now().Weekday().String()
	expected := time.Now().Format("01/02/06")

	result := GetFormatedDateString(currentWeekdayString, "MM/dd/YY")
	if result != expected {
		t.Errorf("Expected: \"%v\", Current: \"%v\"", expected, result)
	}
}
