package main

import (
	"camera/config"
	"camera/logs"
	"camera/server"
	"fmt"
	"os"
)

func init() {
	logs.InitLogStyle()
	err := config.LoadConfig()
	if err != nil {
		logs.Error("fatal error config file: ", err)
		panic(fmt.Errorf("fatal error config file: ", err))
		return
	}
	// 创建目标文件夹
	if err = os.MkdirAll("video", 0755); err != nil {
		logs.Error("创建目录失败:", err)
		return
	}

}

func main() {
	// record video
	recording := server.NewRecording()
	recording.RecordVideo()
}
