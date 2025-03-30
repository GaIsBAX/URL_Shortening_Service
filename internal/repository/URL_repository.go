package repository

import (
	"URL_shortening/internal/model"
	"fmt"
)

type URLRepository struct {
	urls map[string]model.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{urls: make(map[string]model.URL)}
}

func (ur *URLRepository) AddURL(shortURL string, fullURL string) error {
	ur.urls[shortURL] = model.URL{ShortURL: shortURL, FullURL: fullURL}

	return nil
}

func (ur *URLRepository) GetFullURL(shortURL string) (string, error) {
	url, ok := ur.urls[shortURL]
	if !ok {
		return "", fmt.Errorf("short URL %s not found", shortURL)
	}

	return url.FullURL, nil

}
