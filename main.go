package main

import (
	"camera/config"
	"camera/logs"
	"camera/server"
	"fmt"
)

func init() {
	logs.InitLogStyle()
	err := config.LoadConfig()
	if err != nil {
		logs.Error("fatal error config file: %s", err)
		panic(fmt.Errorf("fatal error config file: %s", err))
		return
	}
}

func main() {
	// record video
	recording := server.NewRecording()
	recording.RecordVideo()
}
