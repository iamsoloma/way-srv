package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Mode string `yaml:"mode" env-default:"prod"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:2210"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`

}

func GetConf() *Config {
	configPath := os.Getenv("way-srv-config")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Can`t read a config!")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Can`t read a " + configPath)
	}
	return &cfg
}