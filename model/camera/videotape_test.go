package camera

import "testing"

func TestVideotape_RecordVideo(t *testing.T) {
	//width int, height int, timeout int, bitrate int
	NewVideotape(640, 480, 5, 5).RecordVideo("test.mp4")
}
