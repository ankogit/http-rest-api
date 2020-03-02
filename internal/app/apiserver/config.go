package apiserver


type Config struct {
	BingAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config  {
	return &Config{
		BingAddr:":8080",
		LogLevel: "debug",
	}
}