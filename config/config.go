/*
@author: ledger
@since: 2024/1/29
*/

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// GConfigs global config
var GConfigs *Configs

// Configs global config
type Configs struct {
	Server ServerConfig `yaml:"server"`
	Camera CameraConfig `yaml:"camera"`
}

// ServerConfig 结构体定义服务器配置
type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type CameraConfig struct {
	Shot  CameraShotConfig  `yaml:"shot"`
	Video CameraVideoConfig `yaml:"video"`
}

// CameraShotConfig 结构体定义相机拍摄配置
type CameraShotConfig struct {
	Width   int `yaml:"width"`
	Height  int `yaml:"height"`
	Timeout int `yaml:"timeout"`
	Bitrate int `yaml:"bitrate"`
}

// CameraVideoConfig 结构体定义相机录制视频配置
type CameraVideoConfig struct {
	Width   int `yaml:"width"`
	Height  int `yaml:"height"`
	Timeout int `yaml:"timeout"`
	Bitrate int `yaml:"bitrate"`
}

// LoadConfig use viper to load config
func LoadConfig() error {
	vconfig := viper.New()
	vconfig.AutomaticEnv()
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	vconfig.SetConfigName("config")
	vconfig.AddConfigPath(".")
	vconfig.SetConfigType("yaml")
	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	if err = vconfig.Unmarshal(&GConfigs); err != nil {
		log.Panicln("unmarshal cng file fail " + err.Error())
	}
	return err
}
