package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/configs"
	"github.com/tnp2004/quickdup/modules/auth/authException"
	"github.com/tnp2004/quickdup/pkg/databases"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type Auth interface {
	Login(c echo.Context) error
}

type authImpl struct {
	cfg *configs.Jwt
	db  databases.Database
}

var (
	refreshTokenCookieName = "refreshToken"
	refreshTokenCookiePath = "/api/v1/auth/refreshtoken"
)

func NewAuth(cfg *configs.Jwt, db databases.Database) Auth {
	return &authImpl{cfg, db}
}

func (a *authImpl) Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := utils.BindRequestBody(c, req); err != nil {
		return utils.MessageResp(c, http.StatusBadRequest, err.Error())
	}
	userLoginData, err := a.queryLoginData(req)
	if err != nil {
		return utils.MessageResp(c, http.StatusUnauthorized, err.Error())
	}

	if err := a.LoginAuthentication(userLoginData); err != nil {
		return utils.MessageResp(c, http.StatusUnauthorized, err.Error())
	}

	accessToken, err := a.generateAccessKey(userLoginData.UserID)
	if err != nil {
		return utils.MessageResp(c, http.StatusUnauthorized, err.Error())
	}
	refreshToken, err := a.generateRefreshKey(userLoginData.UserID)
	if err != nil {
		return utils.MessageResp(c, http.StatusUnauthorized, err.Error())
	}

	refreshTokenCookie := setCookie(refreshTokenCookieName, refreshToken, refreshTokenCookiePath, int(a.cfg.RefreshTokenExpireDuration))
	c.SetCookie(&refreshTokenCookie)

	return utils.CustomResp(c, http.StatusOK, &CredentialsResponse{AccessToken: accessToken})
}

func (a *authImpl) LoginAuthentication(req *LoginAuthentication) error {
	if err := utils.ComparePassword(req.HashPassword, req.Password); err != nil {
		return &authException.EmailOrPasswordInvalid{}
	}

	return nil
}

func (a *authImpl) queryLoginData(req *LoginRequest) (*LoginAuthentication, error) {
	query := "SELECT id,password FROM users WHERE email = $1;"
	resp := new(LoginAuthentication)
	args := utils.MakeArgs(req.Email)

	if err := a.db.QueryRowTransaction(query, args, &resp.UserID, &resp.HashPassword); err != nil {
		return nil, err
	}
	resp.Password = req.Password
	return resp, nil
}

func setCookie(name, value, path string, maxAge int) http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		MaxAge:   maxAge,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	return cookie
}
