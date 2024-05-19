package util

import "testing"

func TestGetTodayDateString(t *testing.T) {
	t.Log(GetTodayDateString())
	t.Log(GetTodayHourString())
	t.Log(GetTodayMinuteString())
	t.Log(GetTodaySecondString())
}
