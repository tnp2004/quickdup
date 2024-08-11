package usersController

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/modules/models"
	"github.com/tnp2004/quickdup/modules/users/usersUsecase"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type UsersController interface {
	AddNewUser(c echo.Context) error
}

type usersControllerImpl struct {
	usersUsecase usersUsecase.UsersUsecase
}

func NewUsersController(usersUsecase usersUsecase.UsersUsecase) UsersController {
	return &usersControllerImpl{usersUsecase}
}

func (ctrl *usersControllerImpl) AddNewUser(c echo.Context) error {
	req := new(models.UserRegisterRequest)
	if err := utils.BindRequestBody(c, req); err != nil {
		return utils.MessageResp(c, http.StatusBadRequest, "invalid body request")
	}
	if err := ctrl.usersUsecase.AddNewUser(req); err != nil {
		return utils.MessageResp(c, http.StatusInternalServerError, err.Error())
	}

	return utils.MessageResp(c, http.StatusCreated, fmt.Sprintf("user email %s has been created", req.Email))
}
