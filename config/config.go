package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"prod"`
	Mode string `yaml:"mode" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-default:"./blockchains"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"0.0.0.0:1436"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func GetConf() *Config {
	configPath := "./metadata/config.yaml" //os.Getenv("config")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Can`t open a config!\n" + "Config Path is " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Can`t read a " + configPath)
	}
	return &cfg
}
