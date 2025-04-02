package repository

import (
	"URL_shortening/internal/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

type URLRepository struct {
	urls map[string]model.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{urls: make(map[string]model.URL)}
}

func (ur *URLRepository) AddURL(shortURL string, fullURL string) error {

	if _, exists := ur.urls[shortURL]; exists {
		return gin.Error{Err: fmt.Errorf("short URL %s already exists", shortURL), Type: gin.ErrorTypePublic}
	}

	ur.urls[shortURL] = model.URL{ShortURL: shortURL, FullURL: fullURL}
	return nil
}

func (ur *URLRepository) GetFullURL(shortURL string) (string, error) {
	url, exists := ur.urls[shortURL]
	if !exists {
		return "", gin.Error{Err: fmt.Errorf("short URL %s not found", shortURL), Type: gin.ErrorTypePublic}
	}

	return url.FullURL, nil
}
