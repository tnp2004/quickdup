package servers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes() {
	_ = s.server

	s.healthRoutes()
}

type healthResp struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func (s *Server) healthRoutes() {
	r := s.server.Group("/health")

	r.GET("/server", func(c echo.Context) error {
		resp := &healthResp{
			Message: "ok",
			Date:    time.Now(),
		}
		return c.JSON(http.StatusOK, resp)
	})

	r.GET("/database", func(c echo.Context) error {
		resp := &healthResp{
			Date: time.Now(),
		}
		if err := s.db.HealthCheck(); err != nil {
			resp.Message = "not ok"
			return c.JSON(http.StatusServiceUnavailable, resp)
		}

		resp.Message = "ok"
		return c.JSON(http.StatusOK, resp)
	})
}
