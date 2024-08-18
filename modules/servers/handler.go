package servers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) registerModuleRouters() {
	s.registerHealthRouter()
	s.registerNotesRouter()
	s.registerUsersRouter()
	s.registerAuthRouter()

	s.server.GET("/api/v1/test", s.authMiddleware.Authorization(func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}))
}
