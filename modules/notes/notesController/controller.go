package notesController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/modules/models"
	"github.com/tnp2004/quickdup/modules/notes/notesUsecase"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type NotesController interface {
	AddNewNote(c echo.Context) error
	GenerateCode(c echo.Context) error
}

type notesControllerImpl struct {
	notesUsecase notesUsecase.NotesUsecase
}

func NewNotesController(notesUsecase notesUsecase.NotesUsecase) NotesController {
	return &notesControllerImpl{notesUsecase}
}

func (ctrl *notesControllerImpl) AddNewNote(c echo.Context) error {
	req := new(models.InsertNoteRequest)
	if err := utils.BindRequestBody(c, req); err != nil {
		return utils.MessageResp(c, http.StatusBadRequest, "invalid body request")
	}
	resp, err := ctrl.notesUsecase.AddNewNote(req)
	if err != nil {
		return utils.MessageResp(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctrl *notesControllerImpl) GenerateCode(c echo.Context) error {
	req := new(models.NoteCode)
	if err := utils.BindRequestBody(c, req); err != nil {
		return utils.MessageResp(c, http.StatusBadRequest, "invalid body request")
	}
	resp, err := ctrl.notesUsecase.GenerateCode(req)
	if err != nil {
		return utils.MessageResp(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}
