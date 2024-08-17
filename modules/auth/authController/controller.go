package authController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/modules/auth/authModels"
	"github.com/tnp2004/quickdup/modules/auth/authUsecase"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type AuthController interface {
	Login(c echo.Context) error
}

type authControllerImpl struct {
	authUsecase authUsecase.AuthUsecase
}

func NewAuthController(authUsecase authUsecase.AuthUsecase) AuthController {
	return &authControllerImpl{authUsecase}
}

func (ctrl *authControllerImpl) Login(c echo.Context) error {
	req := new(authModels.LoginRequest)
	if err := utils.BindRequestBody(c, req); err != nil {
		return utils.MessageResp(c, http.StatusBadRequest, err.Error())
	}

	credentials, err := ctrl.authUsecase.Login(req)
	if err != nil {
		return utils.MessageResp(c, http.StatusUnauthorized, err.Error())
	}
	refreshTokenCookie := ctrl.authUsecase.SetRefreshTokenCookie(credentials.RefreshToken)
	c.SetCookie(&refreshTokenCookie)

	return utils.CustomResp(c, http.StatusOK, &authModels.CredentialsResponse{AccessToken: credentials.AccessToken})
}
