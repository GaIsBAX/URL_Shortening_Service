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

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

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

	// w.Header().Set("Content-Type", "text/plain")
	// w.Header().Set("Location", shortURL)
	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(shortURL))
}

func (h *Handler) RedirectHandler(c *gin.Context) {

	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	shortURL := c.Param("shortURL")

	fullURL, err := h.Service.GetFullURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// w.WriteHeader(http.StatusTemporaryRedirect)
	// http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
	c.Redirect(http.StatusTemporaryRedirect, fullURL)
	// http.Redirect(c.Writer, c.Request, fullURL, http.StatusTemporaryRedirect)
}
