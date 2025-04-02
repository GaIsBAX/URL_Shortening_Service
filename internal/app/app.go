package app

import (
	"URL_shortening/internal/config"
	"URL_shortening/internal/repository"
	"URL_shortening/internal/service"

	"github.com/gin-gonic/gin"
)

func Run() error {

	r := gin.Default()
	cfg := config.InitConfig()

	repo := repository.NewURLRepository()
	service := service.NewURLService(repo)
	handler := &Handler{Service: service}

	r.POST("/", handler.ShortenHandler)
	r.GET("/:shortURL", handler.RedirectHandler)
	return r.Run(cfg.Address)
}
