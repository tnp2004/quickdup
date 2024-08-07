package servers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/configs"
)

type Server struct {
	cfg    *configs.Config
	server *echo.Echo
}

func NewServer(cfg *configs.Config) *Server {
	return &Server{
		cfg:    cfg,
		server: echo.New(),
	}
}

func (s *Server) Start() {
	e := s.server

	s.RegisterRoutes()

	port := fmt.Sprintf(":%s", s.cfg.Server.Port)
	e.Logger.Fatal(e.Start(port))
}
