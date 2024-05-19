package main

import (
	"flag"
	"os"
)

// config структура для хранения параметров запуска
type config struct {
	ServerAddr      string `env:"SERVER_ADDRESS"`
	RootUrl         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

// parseConfig парсит конфигурационные параметры со следующим приоритетом:
// env > cli option
func (c *config) parseConfig() {
	flag.StringVar(&c.ServerAddr, "a", "localhost:8080", "server listen address")
	flag.StringVar(&c.RootUrl, "b", "/", "root url")
	flag.StringVar(&c.FileStoragePath, "f", "/tmp/short-url-db.json", "file storage path")
	flag.StringVar(&c.DatabaseDSN, "d", "", "database DSN")

	flag.Parse()

	if runRootUrl := os.Getenv("BASE_URL"); runRootUrl != "" {
		c.RootUrl = runRootUrl
	}
	if runFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); runFileStoragePath != "" {
		c.FileStoragePath = runFileStoragePath
	}
	if runServerAddress := os.Getenv("SERVER_ADDRESS"); runServerAddress != "" {
		c.ServerAddr = runServerAddress
	}
	if runDataBaseDSN := os.Getenv("DATABASE_DSN"); runDataBaseDSN != "" {
		c.DatabaseDSN = runDataBaseDSN
	}
}
