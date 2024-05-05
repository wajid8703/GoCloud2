
package server

import (
	"handlers"
)

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
}
