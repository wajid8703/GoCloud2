package server

import (
	"GoCloud2/handlers"
	"GoCloud2/model"
	"context"
)

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
	handlers.FrontPage(s.mux)
	handlers.NewsletterSignup(s.mux, &signupperMock{})
	handlers.NewsletterThanks(s.mux)
}

type signupperMock struct{}

func (s signupperMock) SignupForNewsLetter(ctx context.Context, email model.Email) (string, error) {
	return "", nil
}
