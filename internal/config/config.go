package config

import (
	"os"
)

type Config struct {
	Address string
	BaseURL string
}

func getEnvOrDefault(envVar, defaultValue string) string {
	if value, exists := os.LookupEnv(envVar); exists {
		return value
	}
	return defaultValue
}

func InitConfig() *Config {
	cfg := &Config{}

	defaultAddress := ":8080"
	defaultBaseURL := "http://localhost:8080/"

	cfg.Address = getEnvOrDefault("SERVER_ADDRESS", defaultAddress)
	cfg.BaseURL = getEnvOrDefault("BASE_URL", defaultBaseURL)

	return cfg
}
