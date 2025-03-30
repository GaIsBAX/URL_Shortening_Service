package service

import (
	"URL_shortening/internal/repository"
	"fmt"
	"strings"
)

type URLService struct {
	urlRepository *repository.URLRepository
}

func NewURLService(urlRepository *repository.URLRepository) *URLService {
	return &URLService{urlRepository: urlRepository}
}

func (us *URLService) GenerateShortURL(URL string) (string, error) {

	if URL = strings.TrimSpace(URL); URL == "" {
		return "", fmt.Errorf("url cannot be empty string")
	}

	hash := 0
	for _, char := range URL {
		hash = int(char) + (hash << 6) + (hash << 16) - hash
	}
	shortURL := fmt.Sprintf("%x", hash)
	us.urlRepository.AddURL(shortURL, URL)

	return "http://localhost:8080/" + shortURL, nil
}

func (us *URLService) GetFullURL(shortURL string) (string, error) {
	return us.urlRepository.GetFullURL(shortURL)
}
