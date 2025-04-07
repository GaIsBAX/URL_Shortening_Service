package app

import (
	"URL_shortening/internal/config"
	"URL_shortening/internal/repository"
	"URL_shortening/internal/service"
	"URL_shortening/internal/transport/http"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() error {

	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error loading .env file %s %s \n%s", time.Now().Format("2006/01/02 15:04:05"), time.Now().Format("2006/01/02 15:04:05"), err))
	}

	cfg := config.InitConfig()
	fmt.Printf("Конфигурация загружена: адрес сервера - %s, базовый URL - %s\n", cfg.Address, cfg.BaseURL)

	repo := repository.NewInMemoryRepository()

	urlService := service.NewURLService(repo, cfg) // (repo, cfg)

	handler := http.NewURLHandler(urlService)

	r := gin.Default()
	r.POST("/", handler.ShortenHandler)
	r.GET("/:shortURL", handler.RedirectHandler)

	// r.SetTrustedProxies(nil)
	return r.Run(cfg.Address)
}
