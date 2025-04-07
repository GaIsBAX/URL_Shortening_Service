package repository

import (
	"URL_shortening/internal/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

type URLRepository interface {
	Save(shortURL string, fullURL string) error
	Get(shortURL string) (string, error)
}

type inMemoryRepo struct {
	urls map[string]model.URL
}

func NewInMemoryRepository() URLRepository {
	return &inMemoryRepo{urls: make(map[string]model.URL)}
}

func (r *inMemoryRepo) Save(shortURL string, fullURL string) error {

	if _, exists := r.urls[shortURL]; exists {
		return gin.Error{Err: fmt.Errorf("short URL %s already exists", shortURL), Type: gin.ErrorTypePublic}
	}

	r.urls[shortURL] = model.URL{ShortURL: shortURL, FullURL: fullURL}
	return nil
}

func (r *inMemoryRepo) Get(shortURL string) (string, error) {
	url, exists := r.urls[shortURL]
	if !exists {
		return "", gin.Error{Err: fmt.Errorf("short URL %s not found", shortURL), Type: gin.ErrorTypePublic}
	}

	return url.FullURL, nil
}
