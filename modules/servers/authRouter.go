package servers

import "github.com/tnp2004/quickdup/modules/auth"

func (s *Server) registerAuthRouter() {
	r := s.server.Group("/api/v1/auth")

	auth := auth.NewAuth(s.cfg.Auth.Jwt, s.db)

	r.POST("/login", auth.Login)
}
