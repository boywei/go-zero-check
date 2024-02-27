package main

import (
	"flag"
	"fmt"
	"github.com/boywei/go-zero-check/config"
	"github.com/boywei/go-zero-check/internal/routers"
	"github.com/davecgh/go-spew/spew"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "./config/app.yaml", "config file")
}

func main() {
	flag.Parse()
	config.SetConfig(configFile)
	conf := config.GetConfig()
	spew.Dump(conf)
	r := routers.Router()
	_ = r.Run(conf.Http.Host + ":" + fmt.Sprint(conf.Http.Port)) // 监听并在 0.0.0.0:8080 上启动服务
}
