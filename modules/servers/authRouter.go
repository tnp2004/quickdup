package servers

import (
	"github.com/tnp2004/quickdup/modules/auth/authController"
	"github.com/tnp2004/quickdup/modules/auth/authRepository"
	"github.com/tnp2004/quickdup/modules/auth/authUsecase"
)

func (s *Server) registerAuthRouter() {
	r := s.server.Group("/api/v1/auth")

	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecase(authRepository, s.cfg.Auth.Jwt)
	authController := authController.NewAuthController(authUsecase)

	r.POST("/login", authController.Login)
	r.POST("/token/revoke", authController.RevokeToken)
}
