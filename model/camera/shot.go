package camera

import (
	"camera/logs"
	"fmt"
	"os/exec"
)

type Shot struct {
	MainCommand      string
	OParameter       string
	WParameter       string
	HParameter       string
	TimeoutParameter string
	PicName          string
	Width            int
	Height           int
	Timeout          int
}

func NewShot(picName string, width int, height int, timeout int) *Shot {
	return &Shot{
		MainCommand:      "libcamera-jpeg",
		OParameter:       "-o",
		WParameter:       "--width",
		HParameter:       "--height",
		TimeoutParameter: "--timeout",
		PicName:          picName,
		Width:            width,
		Height:           height,
		Timeout:          timeout,
	}
}

func (s *Shot) Photograph() {
	cmd := exec.Command(s.MainCommand, s.OParameter, s.PicName,
		s.WParameter, fmt.Sprintf("%d", s.Width),
		s.HParameter, fmt.Sprintf("%d", s.Height),
		s.TimeoutParameter, fmt.Sprintf("%d", s.Timeout))

	err := cmd.Run()
	if err != nil {
		logs.Error("error to exec command!!: %s\n", err)
	}
}
