package servers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/modules/auth/authMiddleware"
)

func (s *Server) registerModuleRouters() {
	s.registerHealthRouter()
	s.registerNotesRouter()
	s.registerUsersRouter()
	s.registerAuthRouter()

	s.server.GET("/api/v1/test", authMiddleware.Authorization(s.db, func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}))
}
