package logs

import (
	"log"
)

// 定义不同的日志级别
const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var logPrefixes = map[int]string{
	DebugLevel: "[DEBUG]",
	InfoLevel:  "[INFO]",
	WarnLevel:  "[WARN]",
	ErrorLevel: "[ERROR]",
}

// 设置日志前缀
func setLogPrefix(level int) {
	prefix, ok := logPrefixes[level]
	if !ok {
		prefix = "[UNKNOWN]"
	}
	log.SetPrefix(prefix + " ")
}

// InitLogStyle 初始化日志格式
func InitLogStyle() {
	log.SetPrefix("[go-raspberry-camera-logs] ")
	log.SetFlags(log.LstdFlags)
}

// Debug 打印调试信息
func Debug(a ...any) {
	setLogPrefix(DebugLevel)
	log.Println(a...)
}

// Info 打印一般信息
func Info(a ...any) {
	setLogPrefix(InfoLevel)
	log.Println(a...)
}

// Warn 打印警告信息
func Warn(a ...any) {
	setLogPrefix(WarnLevel)
	log.Println(a...)
}

// Error 打印错误信息
func Error(a ...any) {
	setLogPrefix(ErrorLevel)
	log.Println(a...)
}
