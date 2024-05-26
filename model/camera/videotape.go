package camera

import (
	"camera/logs"
	"fmt"
	"os/exec"
)

type Videotape struct {
	MainCommand      string
	OParameter       string
	WParameter       string
	HParameter       string
	BitrateParameter string
	TimeoutParameter string
	Width            int
	Height           int
	Bitrate          int
	Timeout          int
}

func NewVideotape(width int, height int, timeout int, bitrate int) *Videotape {
	return &Videotape{
		MainCommand:      "libcamera-vid",
		OParameter:       "-o",
		WParameter:       "--width",
		HParameter:       "--height",
		BitrateParameter: "--bitrate",
		TimeoutParameter: "--timeout",
		Width:            width,
		Height:           height,
		Bitrate:          bitrate,
		Timeout:          timeout,
	}
}

func (v *Videotape) RecordVideo(outputFile string) {
	command := []string{v.MainCommand, v.OParameter, outputFile,
		v.WParameter, fmt.Sprintf("%d", v.Width),
		v.HParameter, fmt.Sprintf("%d", v.Height),
		v.BitrateParameter, fmt.Sprintf("%d", v.Bitrate),
		v.TimeoutParameter, fmt.Sprintf("%d", v.Timeout)}

	fmt.Println("指令集是:", command)

	cmd := exec.Command(command[0], command[1:]...)
	err := cmd.Run()
	if err != nil {
		logs.Error("命令执行失败: %v \n", err)
	}
}
