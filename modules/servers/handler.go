package servers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/modules/notes/notesController"
	"github.com/tnp2004/quickdup/modules/notes/notesRepository"
	"github.com/tnp2004/quickdup/modules/notes/notesUsecase"
)

func (s *Server) RegisterRoutes() {
	api := s.server.Group("/api")
	v1 := api.Group("/v1")

	s.healthRoutes(v1)

	notesRepository := notesRepository.NewNotesRepositories(s.db)
	notesUsecase := notesUsecase.NewNotesUsecase(notesRepository)
	notesController := notesController.NewNotesController(notesUsecase, v1)
	notesController.RegisterRoutes()
}

type healthResp struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func (s *Server) healthRoutes(route *echo.Group) {
	r := route.Group("/health")

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
