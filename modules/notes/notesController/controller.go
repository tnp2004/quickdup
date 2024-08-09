package notesController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type NotesController interface {
	RegisterRoutes()
}

type notesControllerImpl struct {
	route *echo.Group
}

func NewNotesController(route *echo.Group) NotesController {
	return &notesControllerImpl{route}
}

func (c *notesControllerImpl) RegisterRoutes() {
	r := c.route.Group("/notes")

	r.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from notes")
	})
}
