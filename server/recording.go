package server

import (
	"camera/config"
	"camera/model/camera"
	"camera/model/helper"
	"fmt"
)

type Recording struct {
	c chan int
}

func NewRecording() *Recording {
	return &Recording{
		c: make(chan int, 1),
	}
}

func (r *Recording) RecordVideo() {
	i := 1

	v := camera.NewVideotape(
		config.GConfigs.Camera.Video.Width,
		config.GConfigs.Camera.Video.Height,
		config.GConfigs.Camera.Video.Timeout,
		config.GConfigs.Camera.Video.Bitrate)
	vc := helper.NewVideoConverter()

	go func() {
		index := <-r.c
		vc.ConvertToMP4(fmt.Sprintf("video%d.h264", index))
	}()

	for {
		fmt.Printf("开始录制第%d视频\n", i)
		v.RecordVideo(fmt.Sprintf("video%d.h264", i))
		r.c <- i
		i++
	}
}
