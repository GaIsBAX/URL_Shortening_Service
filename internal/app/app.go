package app

import (
	"URL_shortening/internal/repository"
	"URL_shortening/internal/service"
	"net/http"
)

func Run() error {

	repo := repository.NewURLRepository()
	service := service.NewURLService(repo)
	handler := &Handler{Service: service}

	http.HandleFunc("/", handler.ShortenHandler)
	http.HandleFunc("/{shortURL}", handler.RedirectHandler)
	return http.ListenAndServe(":8080", nil)
}

// func webhook(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	_, _ = w.Write([]byte(`
// 	{
// 	  "response": {
// 		"text": "Извините, я пока ничего не умею"
// 	  },
// 	  "version": "1.0"
// 	}
//   `))
// }
