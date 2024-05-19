package helper

import (
	"camera/logs"
	"camera/util"
	"os"
	"os/exec"
	"path/filepath"
)

type VideoConverter struct{}

func NewVideoConverter() *VideoConverter {
	return &VideoConverter{}
}

func (vc *VideoConverter) ConvertToMP4(inputFile string) {

	path := filepath.Join(
		"video",
		util.GetTodayDateString(),
		util.GetTodayHourString(),
		util.GetTodayMinuteString(),
	)
	// 创建目标文件夹
	if err := os.MkdirAll(path, 0755); err != nil {
		logs.Error("创建目录失败:", err)
		return
	}

	// 构建输出文件路径
	outputFile := filepath.Join(path, util.GetTodaySecondString())
	logs.Debug("执行的脚本文件是:", "ffmpeg", "-f", "h264", "-i", inputFile, "-vcodec", "copy", outputFile)
	// 构建ffmpeg命令来转换视频格式
	cmd := exec.Command("ffmpeg", "-f", "h264", "-i", inputFile, "-vcodec", "copy", outputFile+".mp4")
	if err := cmd.Run(); err != nil {
		logs.Error("ffmpeg执行出错:", err)
		return

	}

	logs.Info("转换成功: %s 到 %s\n", inputFile, outputFile)

	// 删除原视频文件
	if err := os.Remove(inputFile); err != nil {
		logs.Error("删除文件失败:", err)
	} else {
		logs.Info("删除文件成功: %s\n", inputFile)
	}
}
