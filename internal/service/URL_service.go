package service

import (
	"URL_shortening/internal/config"
	"URL_shortening/internal/repository"
	"fmt"
	"strings"
)

type URLService interface {
	GenerateShortURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

type urlService struct {
	repo repository.URLRepository
	cfg  *config.Config
}

func NewURLService(repo repository.URLRepository, config *config.Config) URLService {
	return &urlService{repo: repo, cfg: config}
}

func (s *urlService) GenerateShortURL(fullURL string) (string, error) {

	BaseURL := s.cfg.BaseURL

	if fullURL = strings.TrimSpace(fullURL); fullURL == "" {
		return "", fmt.Errorf("url cannot be empty string")
	}

	hash := 0
	for _, char := range fullURL {
		hash = int(char) + (hash << 6) + (hash << 16) - hash
	}
	shortURL := fmt.Sprintf("%x", hash)
	s.repo.Save(shortURL, fullURL)

	return BaseURL + shortURL, nil
}

func (s *urlService) GetFullURL(shortURL string) (string, error) {
	return s.repo.Get(shortURL)
}
