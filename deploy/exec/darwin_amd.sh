#!/bin/bash

# 设置系统架构
go env -w GOOS=darwin
go env -w GOARCH=amd64

# 打包
go build -o check-backend main.go