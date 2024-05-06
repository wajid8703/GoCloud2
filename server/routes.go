package server

import (
	"GoCloud2/handlers"
)

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
	handlers.FrontPage(s.mux)
}
