package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v11"
)

// config структура для хранения параметров запуска
type config struct {
	ServerAddr string `env:"SERVER_ADDRESS"`
	RootUrl    string `env:"BASE_URL"`
}

// parseConfig парсит конфигурационные параметры со следующим приоритетом:
// env > cli option
func getConfig() (*config, error) {
	var cfg config
	var requireParse = false

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg)

	if cfg.ServerAddr == "" {
		requireParse = true
		flag.StringVar(&cfg.ServerAddr, "a", ":8080", "server listen address")
	}
	if cfg.RootUrl == "" {
		requireParse = true
		flag.StringVar(&cfg.RootUrl, "b", "/", "root url")
	}
	if requireParse {
		flag.Parse()
	}

	return &cfg, nil
}
