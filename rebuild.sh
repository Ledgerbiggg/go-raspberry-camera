#!/bin/bash

git pull

set GOOS linux
set GOARCH arm64

go build -o raspberry-camera

chmod +x raspberry-camera

./raspberry-camera
