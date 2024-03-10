package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"time"
)

type Config struct {
	Env        string     `yaml:"env"`
	HttpServer HttpServer `yaml:"httpserver"`
}

type HttpServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

var once sync.Once
var config Config

func NewConfig(configPath string) *Config {
	once.Do(func() {
		err := cleanenv.ReadConfig(configPath, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
