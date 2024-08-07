package servers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes() {
	app := s.server

	app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
