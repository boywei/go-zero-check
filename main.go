package main

import (
	"flag"
	"fmt"

	"github.com/boywei/go-zero-check/config"
	"github.com/boywei/go-zero-check/internal/middleware"
	"github.com/boywei/go-zero-check/internal/routers"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

var (
	configFile string
	conf       *config.Config
)

func init() {
	flag.StringVar(&configFile, "config", "./config/app.yaml", "config file")
	flag.Parse()

	config.SetConfig(configFile)
	conf = config.GetConfig()
	spew.Dump(conf)
	// 使用redis作为缓存
	err := middleware.SetupRedisDb(conf.Redis.Addr, conf.Redis.Password)
	if err != nil {
		log.Fatalf("init.SetupRedisDb err: %v", err)
	}

}

func main() {
	r := routers.Router()
	_ = r.Run(conf.Http.Host + ":" + fmt.Sprint(conf.Http.Port)) // 监听并在 0.0.0.0:8080 上启动服务
}
