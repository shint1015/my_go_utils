package utils

import "time"

// GetWeek Get the number of weeks in this year
func GetWeek() int {
	t := time.Now()
	_, w := t.ISOWeek()
	return w
}

// GetYear Get the number of years
func GetYear() int {
	return time.Now().Year()
}

// GetMonth Get the number of months in this year
func GetMonth() int {
	return int(time.Now().Month())
}

// GetDay Get the number of days in this month
func getDay() int {
	return time.Now().Day()
}

// GetMonday Get the number of days in this week's Monday
func GetMonday() int {
	return getDay() - int(time.Now().Weekday())
}
