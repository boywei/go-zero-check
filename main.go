package main

import "github.com/boywei/go-zero-check/internal/api"

func main() {
	r := api.Router()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
