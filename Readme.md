## go-raspberry-camera

<p align="center">
  <img src="https://img.shields.io/github/watchers/Ledgerbiggg/Ledgerbiggg">
  <img src="https://img.shields.io/github/stars/Ledgerbiggg/Ledgerbiggg">
  <img src="https://img.shields.io/github/forks/Ledgerbiggg/Ledgerbiggg">
  <img src="https://img.shields.io/github/issues/Ledgerbiggg/Ledgerbiggg">
  <img src="https://img.shields.io/github/license/Ledgerbiggg/Ledgerbiggg">
  <img src="https://img.shields.io/github/contributors/Ledgerbiggg/Ledgerbiggg">
</p>

### 这个是基于go语言实现的树莓派摄像头的功能

#### 简介

* 使用树莓派的摄像头并一直录制视频

### 前提条件

* 准备一个树莓派(3b及以上)
* 给树莓派安装好一个摄像头
* 在树莓派上配置好环境

1. 更新软件包

```shell
sudo apt install libcamera-vid
```

2. 下载libcamera-apps

```shell
sudo apt install libcamera-apps
```
3. 运行摄像头的hello world,运行成功则表示可以使用
```shell
libcamera-hello
```

* 运行以下脚本命令:

```shell
curl -o load.sh https://raw.githubusercontent.com/Ledgerbiggg/go-raspberry-camera/main/load.sh
chmod +x load.sh
./load.sh
```

* 就会在当前目录下面生成这样的文件(视频的一些配置请参考config.yml文件):
* ![](https://img2.imgtp.com/2024/05/26/1lWJIUZu.png)



