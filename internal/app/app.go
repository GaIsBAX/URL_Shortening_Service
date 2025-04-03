package app

import (
	"URL_shortening/internal/config"
	"URL_shortening/internal/repository"
	"URL_shortening/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() error {

	r := gin.Default()
	cfg := config.InitConfig()
	fmt.Printf("Конфигурация загружена: адрес сервера - %s, базовый URL - %s\n", cfg.Address, cfg.BaseURL)

	repo := repository.NewURLRepository()
	service := service.NewURLService(repo, cfg)
	handler := &Handler{Service: service}

	r.POST("/", handler.ShortenHandler)
	r.GET("/:shortURL", handler.RedirectHandler)

	r.SetTrustedProxies(nil)
	return r.Run(cfg.Address)
}
