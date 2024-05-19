package util

import "time"

func GetTodayDateString() string {
	return time.Now().Format("2006-01-02")
}

func GetTodayHourString() string {
	return time.Now().Format("15")
}

func GetTodayMinuteString() string {
	return time.Now().Format("04")
}

func GetTodaySecondString() string {
	return time.Now().Format("05")
}
