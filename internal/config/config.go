package config

import (
	"flag"
	"fmt"
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

	defaultAddress := "localhost:8080"
	defaultBaseURL := "http://localhost:8080"

	flagAddress := flag.String("a", defaultAddress, "Адрес HTTP-сервера")
	flagBaseURL := flag.String("b", defaultBaseURL, "Базовый URL для сокращенных ссылок")
	flag.Parse()

	cfg.Address = getEnvOrDefault("SERVER_ADDRESS", *flagAddress)
	cfg.BaseURL = getEnvOrDefault("BASE_URL", *flagBaseURL)

	fmt.Println("Конфигурация загружена:")
	fmt.Printf("Адрес сервера: %s\n", cfg.Address)
	fmt.Printf("Базовый URL: %s\n", cfg.BaseURL)

	return cfg
}
