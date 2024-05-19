package camera

import "testing"

func TestShot_Photograph(t *testing.T) {
	NewShot("test.jpg", 640, 480, 5).Photograph()
}
