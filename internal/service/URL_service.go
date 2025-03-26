package service

import (
	"URL_shortening/internal/repository"
	"fmt"
)

type URLService struct {
	urlRepository *repository.URLRepository
}

func NewURLService(urlRepository *repository.URLRepository) *URLService {
	return &URLService{urlRepository: urlRepository}
}

func (us *URLService) GenerateShortURL(URL string) string {
	hash := 0
	for _, char := range URL {
		hash = int(char) + (hash << 6) + (hash << 16) - hash
	}
	shortURL := fmt.Sprintf("%x", hash)
	us.urlRepository.AddURL(shortURL, URL)
	return shortURL
}
