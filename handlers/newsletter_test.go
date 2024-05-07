package handlers

import (
	"GoCloud2/model"
	"context"
	"github.com/go-chi/chi/v5"
	is2 "github.com/matryer/is"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type signupperMock struct {
	email model.Email
}

func (s *signupperMock) SignupForNewsLetter(ctx context.Context, email model.Email) (string, error) {
	s.email = email
	return "", nil
}

func TestNewsletterSignup(t *testing.T) {
	mux := chi.NewMux()
	s := &signupperMock{}
	NewsletterSignup(mux, s)
	t.Run("signs up a valid email address", func(t *testing.T) {
		is := is2.New(t)
		code, _, _ := makePostRequest(mux, "/newsletter/signup", createFromHeader(), strings.NewReader("email=me%40example.com"))
		is.Equal(http.StatusFound, code)
		is.Equal(model.Email("me@example.com"), s.email)
	})

	t.Run("rejects an invalid address", func(t *testing.T) {
		is := is2.New(t)
		code, _, _ := makePostRequest(mux, "/newsletter/signup", createFromHeader(), strings.NewReader("email=notanemail"))
		is.Equal(http.StatusBadRequest, code)
	})

}

func makePostRequest(handler http.Handler, target string, header http.Header, body io.Reader) (int, http.Header, string) {
	req := httptest.NewRequest(http.MethodPost, target, body)
	req.Header = header
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	result := res.Result()
	bodyBytes, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	return result.StatusCode, result.Header, string(bodyBytes)
}

func createFromHeader() http.Header {
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	return header
}
