package helper

import "testing"

func TestVideoConverter_ConvertToMP4(t *testing.T) {
	NewVideoConverter().ConvertToMP4("test.h264", "test")
}
