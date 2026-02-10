package config

import "os"

type Config struct {
	DSN  string
	Host string
	Port string
}

func LoadConfig() *Config {
	return &Config{
		DSN:  os.Getenv("CONNECTION_STRING"),
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}
