package servers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes() {
	_ = s.server

	s.healthRoutes()
}

func (s *Server) healthRoutes() {
	s.server.GET("/health", func(c echo.Context) error {
		var healthResp struct {
			Server   string `json:"server"`
			Database string `json:"database"`
		}

		healthResp.Server = "OK"
		if err := s.db.HealthCheck(); err != nil {
			healthResp.Database = "NOT OK"
		} else {
			healthResp.Database = "OK"
		}

		return c.JSON(http.StatusOK, healthResp)
	})
}
