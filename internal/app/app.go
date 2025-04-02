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

// func webhook(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	_, _ = w.Write([]byte(`
// 	{
// 	  "response": {
// 		"text": "Извините, я пока ничего не умею"
// 	  },
// 	  "version": "1.0"
// 	}
//   `))
// }
