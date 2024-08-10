package notesController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/modules/models"
	"github.com/tnp2004/quickdup/modules/notes/notesUsecase"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type NotesController interface {
	RegisterRoutes()
}

type notesControllerImpl struct {
	NotesUsecase notesUsecase.NotesUsecase
	route        *echo.Group
}

func NewNotesController(NotesUsecase notesUsecase.NotesUsecase, route *echo.Group) NotesController {
	return &notesControllerImpl{NotesUsecase, route}
}

func (ctrl *notesControllerImpl) RegisterRoutes() {
	r := ctrl.route.Group("/notes")

	r.POST("/", func(c echo.Context) error {
		req := new(models.InsertNoteRequest)
		if err := utils.BindRequestBody(c, req); err != nil {
			return utils.MessageResp(c, http.StatusBadRequest, "invalid body request")
		}
		resp, err := ctrl.NotesUsecase.AddNewNote(req)
		if err != nil {
			return utils.MessageResp(c, http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	})
}
