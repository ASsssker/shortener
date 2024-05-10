package main

import (
	"flag"
	"os"
)

// config структура для хранения параметров запуска
type config struct {
	ServerAddr string `env:"SERVER_ADDRESS"`
	RootUrl    string `env:"BASE_URL"`
}

// parseConfig парсит конфигурационные параметры со следующим приоритетом:
// env > cli option
func (c *config) parseConfig() {
	flag.StringVar(&c.ServerAddr, "a", "localhost:8080", "server listen address")
	flag.StringVar(&c.RootUrl, "b", "/", "root url")
	flag.Parse()

	if runServerAddr := os.Getenv("SERVER_ADDRESS"); runServerAddr != "" {
		c.ServerAddr = runServerAddr
	}
	if runRootUrl := os.Getenv("BASE_URL"); runRootUrl != "" {
		c.RootUrl = runRootUrl
	}
}
