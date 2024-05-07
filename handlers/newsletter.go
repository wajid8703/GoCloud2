package handlers

import (
	"GoCloud2/model"
	"GoCloud2/views"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type signupper interface {
	SignupForNewsLetter(ctx context.Context, email model.Email) (string, error)
}

func NewsletterSignup(mux chi.Router, s signupper) {
	mux.Post("/newsletter/signup", func(w http.ResponseWriter, r *http.Request) {
		email := model.Email(r.FormValue("email"))

		if !email.IsValid() {
			http.Error(w, "email is invalid", http.StatusBadRequest)
			return
		}

		if _, err := s.SignupForNewsLetter(r.Context(), email); err != nil {
			http.Error(w, "error signing up", http.StatusBadGateway)
			return
		}

		http.Redirect(w, r, "/newsletter/thanks", http.StatusFound)
	})
}

func NewsletterThanks(mux chi.Router) {
	mux.Get("/newsletter/thanks", func(w http.ResponseWriter, r *http.Request) {
		_ = views.NewsletterThanksPage("/newsletter/thanks").Render(w)
	})
}
