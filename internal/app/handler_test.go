package app

import (
	"URL_shortening/internal/repository"
	"URL_shortening/internal/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
				Service: service.NewURLService(repository.NewURLRepository()),
			},
		},
		{
			name: "empty body test",
			body: "",
			want: want{
				statusCode: http.StatusNotFound,
			},
			fields: fields{
				Service: service.NewURLService(repository.NewURLRepository()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			h := &Handler{
				Service: tt.fields.Service,
			}
			h.ShortenHandler(w, r)

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
		response    string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "simple test #1",
			want: want{
				contentType: "text/html; charset=utf-8",
				statusCode:  http.StatusTemporaryRedirect,
				response:    "<a href=\"https://google.com\">Temporary Redirect</a>.",
			},
			fields: fields{
				Service: service.NewURLService(repository.NewURLRepository()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Отправляем POST-запрос для сокращения URL
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://google.com"))
			w := httptest.NewRecorder()

			h := &Handler{Service: tt.fields.Service}
			h.ShortenHandler(w, r)

			// 2. Проверяем, что сокращение прошло успешно
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, http.StatusCreated, res.StatusCode)

			shortURL, _ := io.ReadAll(res.Body)
			shortURLStr := strings.TrimSpace(string(shortURL)) // Убираем лишние пробелы

			// 3. Теперь используем этот URL в GET-запросе на редирект
			r = httptest.NewRequest(http.MethodGet, shortURLStr, nil)
			w = httptest.NewRecorder()

			h.RedirectHandler(w, r)

			// 4. Проверяем редирект
			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"))
			assert.Equal(t, tt.want.statusCode, w.Code)

			res = w.Result()
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, tt.want.response, strings.TrimSpace(string(body)))
		})
	}
}
