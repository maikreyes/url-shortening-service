package config

import "os"

type Config struct {
	DSN   string
	Host  string
	Port  string
	Table string
}

func LoadConfig() *Config {
	return &Config{
		DSN:   os.Getenv("CONNECTION_STRING"),
		Host:  os.Getenv("HOST"),
		Port:  os.Getenv("PORT"),
		Table: os.Getenv("TABLE_NAME"),
	}
}
