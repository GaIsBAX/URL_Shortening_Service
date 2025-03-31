package app

import (
	"URL_shortening/internal/service"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Handler struct {
	Service *service.URLService
}

func (h *Handler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	originalURL, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := url.ParseRequestURI(string(originalURL)); err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}

	_, err = client.Get(string(originalURL))
	resp, err := client.Get(string(originalURL))
	if err != nil || resp.StatusCode >= 400 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	shortURL, err := h.Service.GenerateShortURL(string(originalURL))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", shortURL)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[1:]

	fullURL, err := h.Service.GetFullURL(shortURL)
	if err != nil {
		http.Error(w, "URL не найден", http.StatusNotFound)
		return
	}

	// w.WriteHeader(http.StatusTemporaryRedirect)
	http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)

}
