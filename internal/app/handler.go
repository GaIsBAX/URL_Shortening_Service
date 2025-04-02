package app

import (
	"URL_shortening/internal/service"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.URLService
}

func (h *Handler) ShortenHandler(c *gin.Context) {

	originalURL, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if _, err := url.ParseRequestURI(string(originalURL)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(strings.TrimSpace(string(originalURL)))

	if err != nil || resp.StatusCode >= 400 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}
	defer resp.Body.Close()

	shortURL, err := h.Service.GenerateShortURL(string(originalURL))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate short URL"})
		return
	}

	c.Header("Content-Type", "text/plain")
	c.Header("Location", shortURL)
	c.Status(http.StatusCreated)
	c.Writer.Write([]byte(shortURL))
}

func (h *Handler) RedirectHandler(c *gin.Context) {

	shortURL := c.Param("shortURL")
	if shortURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short URL is required"})
		return
	}

	fullURL, err := h.Service.GetFullURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, fullURL)
}
