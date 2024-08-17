package authUsecase

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tnp2004/quickdup/configs"
	"github.com/tnp2004/quickdup/modules/auth/authException"
	"github.com/tnp2004/quickdup/modules/auth/authModels"
	"github.com/tnp2004/quickdup/modules/auth/authRepository"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type AuthUsecase interface {
	Login(req *authModels.LoginRequest) (*authModels.LoginResponse, error)
	SetRefreshTokenCookie(refreshToken string) http.Cookie
}

type authUsecaseImpl struct {
	authRepository authRepository.AuthRository
	jwtCfg         *configs.Jwt
}

func NewAuthUsecase(authRepository authRepository.AuthRository, jwtCfg *configs.Jwt) AuthUsecase {
	return &authUsecaseImpl{authRepository, jwtCfg}
}

var (
	issuer                 = "quickdup"
	refreshTokenCookieName = "rft"
	refreshTokenCookiePath = "/api/v1/auth/token/revoke"
)

func (u *authUsecaseImpl) Login(req *authModels.LoginRequest) (*authModels.LoginResponse, error) {
	userLoginData, err := u.authRepository.QueryLoginData(req)
	if err != nil {
		return nil, &authException.Unauthorized{}
	}

	if err := u.authentication(userLoginData); err != nil {
		return nil, &authException.Unauthorized{}
	}

	accessToken, err := u.generateAccessKey(userLoginData.UserID)
	if err != nil {
		return nil, &authException.Unauthorized{}
	}
	refreshToken, err := u.generateRefreshKey(userLoginData.UserID)
	if err != nil {
		return nil, &authException.Unauthorized{}
	}

	if err := u.authRepository.InsertAuthorizationCredentials(&authModels.AuthorizationCredentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, &authException.Unauthorized{}
	}

	return &authModels.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *authUsecaseImpl) SetRefreshTokenCookie(refreshToken string) http.Cookie {
	return utils.SetCookie(refreshTokenCookieName, refreshToken, refreshTokenCookiePath, int(u.jwtCfg.RefreshTokenExpireDuration))
}

func (u *authUsecaseImpl) authentication(req *authModels.Authentication) error {
	if err := utils.ComparePassword(req.HashPassword, req.Password); err != nil {
		return &authException.EmailOrPasswordInvalid{}
	}

	return nil
}

func (u *authUsecaseImpl) generateAccessKey(userID string) (string, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(u.jwtCfg.AccessTokenExpireDuration))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   userID,
		Audience:  []string{"user"},
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})
	ss, err := token.SignedString([]byte(u.jwtCfg.AccessSecretKey))
	if err != nil {
		log.Printf("error sign access token. Error: %s", err.Error())
		return "", &authException.SignAccessToken{}
	}

	return ss, nil
}

func (u *authUsecaseImpl) generateRefreshKey(userID string) (string, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(u.jwtCfg.RefreshTokenExpireDuration))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   userID,
		Audience:  []string{"user"},
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})
	ss, err := token.SignedString([]byte(u.jwtCfg.RefreshSecretKey))
	if err != nil {
		log.Printf("error sign refresh token. Error: %s", err.Error())
		return "", &authException.SignRefreshToken{}
	}

	return ss, nil
}
