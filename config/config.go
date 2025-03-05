package config

import (
	"log"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

var (
	Once sync.Once
	cfg  *Config
)

func Get() *Config {
	if cfg == nil {
		log.Panic("Configuration has not yet been initialized")
	}
	return cfg
}

type Config struct {
	Server   Server   `env:"server"`
	Domain   Domain   `env:"domain"`
	Postgres Postgres `env:"postgres"`
}

type Server struct {
	Port int `env:"port"`
}

type Domain struct {
	Link string `env:"link"`
}

type Postgres struct {
	Host     string `env:"host"`
	Port     int    `env:"port"`
	User     string `env:"user"`
	Password string `env:"password"`
	Database string `env:"database"`
}

func Environments() {
	Once.Do(func() {
		cfg = new(Config)
		if err := config.GetConfigFromEnv(cfg); err != nil {
			log.Panicf("error parsing enviroment vars \n%v", err)
		}
	})
}
