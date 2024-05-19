package server

import (
	"camera/config"
	"camera/model/camera"
	"camera/model/helper"
	"fmt"
	"time"
)

type Recording struct{}

func NewRecording() *Recording {
	return &Recording{}
}

func (r *Recording) RecordVideo() {
	i := 1

	v := camera.NewVideotape(
		config.GConfigs.Camera.Video.Width,
		config.GConfigs.Camera.Video.Height,
		config.GConfigs.Camera.Video.Timeout,
		config.GConfigs.Camera.Video.Bitrate)
	vc := helper.NewVideoConverter()

	for {
		fmt.Printf("开始录制第%d视频\n", i)
		v.RecordVideo(fmt.Sprintf("video%d.h264", i))
		time.Sleep(1 * time.Second)
		go func(i int) {
			vc.ConvertToMP4(fmt.Sprintf("video%d.h264", i), "output")
		}(i)
		i++
	}
}
