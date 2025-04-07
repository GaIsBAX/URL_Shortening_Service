package http

import (
	"URL_shortening/internal/config"
	"URL_shortening/internal/repository"
	"URL_shortening/internal/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ShortenHandler(t *testing.T) {
	type fields struct {
		Service *service.URLService
	}
	type want struct {
		contentType string
		statusCode  int
		response    string
	}

	tests := []struct {
		name   string
		fields fields
		body   string
		want   want
	}{
		{
			name: "simple test #1",
			body: "https://google.com",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusCreated,
				response:    "http://localhost:8080/-3f0b1fcd459b6745",
			},
			fields: fields{
				// Service: service.NewURLService(repository.NewInMemoryRepository()),
			},
		},
		{
			name: "simple test #2",
			body: "https://http.cat/ ",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusCreated,
				response:    "http://localhost:8080/-59551051e9dbca70",
			},
			fields: fields{
				// Service: service.NewURLService(repository.NewURLRepository(), config.InitConfig()),
			},
		},
		{
			name: "empty body test",
			body: "",
			want: want{
				contentType: "application/json; charset=utf-8",
				statusCode:  http.StatusBadRequest,
				response:    `{"error":"Invalid URL"}`,
			},
			fields: fields{
				// Service: service.NewURLService(repository.NewURLRepository(), config.InitConfig()),
			},
		},
		{
			name: "invalid url test",
			body: "http://goologle.com",
			want: want{
				contentType: "application/json; charset=utf-8",
				statusCode:  http.StatusBadRequest,
				response:    `{"error":"Invalid URL"}`,
			},
			fields: fields{
				// Service: service.NewURLService(repository.NewURLRepository(), config.InitConfig()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewURLHandler(service.NewURLService(repository.NewInMemoryRepository(), config.InitConfig()))

			r.POST("/", h.ShortenHandler)
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"))
			assert.Equal(t, tt.want.statusCode, w.Code)

			res := w.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.want.response, strings.TrimSpace(string(body)))

		})
	}
}

func TestHandler_RedirectHandler(t *testing.T) {
	type fields struct {
		Service *service.URLService
	}

	type want struct {
		contentType string
		statusCode  int
		location    string
	}

	tests := []struct {
		name    string
		request string
		fields  fields
		want    want
	}{
		{
			name:    "simple test #1",
			request: "https://google.com",
			want: want{
				contentType: "text/html; charset=utf-8",
				statusCode:  http.StatusTemporaryRedirect,
				location:    "https://google.com",
			},
			fields: fields{
				// Service: service.NewURLService(repository.NewURLRepository(), config.InitConfig()),
			},
		},
		{
			name:    "simple test #2",
			request: "https://ya.ru",
			want: want{
				contentType: "text/html; charset=utf-8",
				statusCode:  http.StatusTemporaryRedirect,
				location:    "https://ya.ru",
			},
			fields: fields{
				// Service: service.NewURLService(repository.NewURLRepository(), config.InitConfig()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()

			h := NewURLHandler(service.NewURLService(repository.NewInMemoryRepository(), config.InitConfig()))

			r.POST("/", h.ShortenHandler)
			r.GET("/:shortURL", h.RedirectHandler)

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.request))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusCreated, w.Code)
			shortURL, _ := io.ReadAll(w.Body)
			shortURLStr := strings.TrimSpace(string(shortURL))

			req = httptest.NewRequest(http.MethodGet, shortURLStr, nil)
			w = httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"))
			assert.Equal(t, tt.want.statusCode, w.Code)
			assert.Equal(t, tt.want.location, w.Header().Get("Location"))

		})
	}
}
