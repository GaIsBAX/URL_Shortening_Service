package app

import (
	"URL_shortening/internal/service"
	"io"
	"net/http"
)

type Handler struct {
	Service *service.URLService
}

func (h *Handler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	url, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL, err := h.Service.GenerateShortURL(string(url))

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
