package config

import "os"

type Config struct {
	DSN         string
	Driver      string
	Host        string
	Port        string
	UrlTable    string
	UserTable   string
	Environment string
}

func LoadConfig() *Config {
	return &Config{
		DSN:         os.Getenv("CONNECTION_STRING"),
		Driver:      os.Getenv("DB_DRIVER"),
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		UrlTable:    os.Getenv("URL_TABLE_NAME"),
		UserTable:   os.Getenv("USER_TABLE_NAME"),
		Environment: os.Getenv("ENVIRONMENT"),
	}
}
