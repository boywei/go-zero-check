package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Http  HttpConfig  `yaml:"http"`
	Grpc  GrpcConfig  `yaml:"grpc"`
	Redis RedisConfig `yaml:"redis"`
}

type HttpConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type GrpcConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

var (
	config *Config
)

func SetConfig(configFile string) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Fatal error config file: %s \n", err)
	}
	config = &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Errorf("Fatal error config file: %s \n", err)
	}
}

func GetConfig() *Config {
	return config
}
