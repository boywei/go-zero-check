#!/bin/bash

# 设置系统架构
go env -w GOOS=windows
go env -w GOARCH=arm64

# 打包
go build -o check-backend main.go