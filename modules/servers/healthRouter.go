package servers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/modules/auth/authMiddleware"
)

type healthResp struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func (s *Server) registerHealthRouter() {
	r := s.server.Group("/api/v1/health")

	r.GET("/server", authMiddleware.Authorization(s.db, s.serverHealth))
	r.GET("/database", authMiddleware.Authorization(s.db, s.databaseHealth))

}

func (s *Server) serverHealth(c echo.Context) error {
	resp := &healthResp{
		Message: "ok",
		Date:    time.Now(),
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) databaseHealth(c echo.Context) error {
	resp := &healthResp{
		Date: time.Now(),
	}
	if err := s.db.HealthCheck(); err != nil {
		resp.Message = "not ok"
		return c.JSON(http.StatusServiceUnavailable, resp)
	}

	resp.Message = "ok"
	return c.JSON(http.StatusOK, resp)
}
