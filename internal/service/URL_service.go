package service

import (
	"URL_shortening/internal/config"
	"URL_shortening/internal/repository"
	"fmt"
	"strings"
)

type URLService struct {
	urlRepository *repository.URLRepository
	cfg           *config.Config
}

func NewURLService(urlRepository *repository.URLRepository, cfg *config.Config) *URLService {
	return &URLService{urlRepository: urlRepository, cfg: cfg}
}

func (us *URLService) GenerateShortURL(URL string) (string, error) {

	BaseURL := us.cfg.BaseURL

	if URL = strings.TrimSpace(URL); URL == "" {
		return "", fmt.Errorf("url cannot be empty string")
	}

	hash := 0
	for _, char := range URL {
		hash = int(char) + (hash << 6) + (hash << 16) - hash
	}
	shortURL := fmt.Sprintf("%x", hash)
	us.urlRepository.AddURL(shortURL, URL)

	return BaseURL + shortURL, nil
}

func (us *URLService) GetFullURL(shortURL string) (string, error) {
	return us.urlRepository.GetFullURL(shortURL)
}
