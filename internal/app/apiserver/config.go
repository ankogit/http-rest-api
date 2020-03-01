package apiserver

import "github.com/ankogit/http-rest-api/internal/app/store"

type Config struct {
	BingAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	Store *store.Config
}

func NewConfig() *Config  {
	return &Config{
		BingAddr:":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}