package helper

import (
	"camera/logs"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type VideoConverter struct{}

func NewVideoConverter() *VideoConverter {
	return &VideoConverter{}
}

func (vc *VideoConverter) ConvertToMP4(inputFile, outputDir string) {
	todayDateString := getTodayDateString()
	outputPath := filepath.Join(outputDir, todayDateString)

	// 创建目标文件夹
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		logs.Error("创建目录失败:", err)
		return
	}

	// 构建输出文件路径
	outputFile := filepath.Join(outputPath, filepath.Base(inputFile))

	// 构建ffmpeg命令来转换视频格式
	cmd := exec.Command("ffmpeg", "-f", "h264", "-i", inputFile, "-vcodec", "copy", outputFile)
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

func getTodayDateString() string {
	return time.Now().Format("2006-01-02")
}
