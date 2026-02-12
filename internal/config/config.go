package config

import "os"

type Config struct {
	DSN         string
	Driver      string
	Host        string
	Port        string
	Table       string
	Environment string
}

func LoadConfig() *Config {
	return &Config{
		DSN:         os.Getenv("CONNECTION_STRING"),
		Driver:      os.Getenv("DB_DRIVER"),
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		Table:       os.Getenv("TABLE_NAME"),
		Environment: os.Getenv("ENVIRONMENT"),
	}
}
