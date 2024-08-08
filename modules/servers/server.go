package servers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/configs"
	"github.com/tnp2004/quickdup/pkg/databases"
)

type Server struct {
	cfg    *configs.Config
	server *echo.Echo
	db     databases.Database
}

func NewServer(cfg *configs.Config, db databases.Database) *Server {
	return &Server{
		cfg:    cfg,
		server: echo.New(),
		db:     db,
	}
}

func (s *Server) Start() {
	e := s.server

	s.RegisterRoutes()

	port := fmt.Sprintf(":%s", s.cfg.Server.Port)
	e.Logger.Fatal(e.Start(port))
}
