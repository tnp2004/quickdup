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
	Login(req *authModels.LoginRequest) (*authModels.Credentials, error)
	SetRefreshTokenCookie(refreshToken string) http.Cookie
	RevokeToken(refreshToken string) (*authModels.Credentials, error)
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

func (u *authUsecaseImpl) Login(req *authModels.LoginRequest) (*authModels.Credentials, error) {
	userLoginData, err := u.authRepository.QueryLoginData(req)
	if err != nil {
		return nil, &authException.EmailNotFound{}
	}

	if err := u.authentication(userLoginData); err != nil {
		return nil, err
	}

	accessTokenExpiresAt := time.Now().Add(time.Second * time.Duration(u.jwtCfg.AccessTokenExpireDuration))
	accessToken, err := u.generateAccessKey(userLoginData.UserID, jwt.NewNumericDate(accessTokenExpiresAt))
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().Add(time.Second * time.Duration(u.jwtCfg.RefreshTokenExpireDuration))
	refreshToken, err := u.generateRefreshKey(userLoginData.UserID, jwt.NewNumericDate(expiresAt))
	if err != nil {
		return nil, err
	}

	if err := u.authRepository.InsertAuthorizationCredentials(&authModels.AuthorizationCredentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, err
	}

	return &authModels.Credentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *authUsecaseImpl) SetRefreshTokenCookie(refreshToken string) http.Cookie {
	return utils.SetCookie(refreshTokenCookieName, refreshToken, refreshTokenCookiePath, int(u.jwtCfg.RefreshTokenExpireDuration))
}

func (u *authUsecaseImpl) RevokeToken(refreshToken string) (*authModels.Credentials, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &authException.UnexpectedSigningMethod{}
		}
		return []byte(u.jwtCfg.RefreshSecretKey), nil
	})
	if err != nil {
		log.Printf("error authorization. Error: %s", err.Error())
		return nil, err
	}

	if !token.Valid {
		log.Println("error invalid token")
		return nil, err
	}

	if err := u.authRepository.DeleteCredential(refreshToken); err != nil {
		return nil, err
	}

	userID, err := token.Claims.GetSubject()
	if err != nil {
		log.Printf("error get user id from token. Error: %s", err.Error())
		return nil, &authException.RevokeToken{}
	}
	refreshTokenExpiresAt, err := token.Claims.GetExpirationTime()
	if err != nil {
		log.Printf("error get expiration time from token. Error: %s", err.Error())
		return nil, &authException.RevokeToken{}
	}
	accessTokenExpiresAt := time.Now().Add(time.Second * time.Duration(u.jwtCfg.AccessTokenExpireDuration))
	newAccessToken, err := u.generateAccessKey(userID, jwt.NewNumericDate(accessTokenExpiresAt))
	if err != nil {
		return nil, &authException.RevokeToken{}
	}
	newRefreshToken, err := u.generateRefreshKey(userID, refreshTokenExpiresAt)
	if err != nil {
		return nil, &authException.RevokeToken{}
	}

	if err := u.authRepository.InsertAuthorizationCredentials(&authModels.AuthorizationCredentials{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}); err != nil {
		return nil, &authException.RevokeToken{}
	}

	return &authModels.Credentials{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (u *authUsecaseImpl) authentication(req *authModels.Authentication) error {
	if err := utils.ComparePassword(req.HashPassword, req.Password); err != nil {
		return &authException.EmailOrPasswordInvalid{}
	}

	return nil
}

func (u *authUsecaseImpl) generateAccessKey(userID string, expiresAt *jwt.NumericDate) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   userID,
		Audience:  []string{"user"},
		ExpiresAt: expiresAt,
		NotBefore: jwt.NewNumericDate(time.Now()),
	})
	ss, err := token.SignedString([]byte(u.jwtCfg.AccessSecretKey))
	if err != nil {
		log.Printf("error sign access token. Error: %s", err.Error())
		return "", &authException.SignAccessToken{}
	}

	return ss, nil
}

func (u *authUsecaseImpl) generateRefreshKey(userID string, expiresAt *jwt.NumericDate) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   userID,
		Audience:  []string{"user"},
		ExpiresAt: expiresAt,
		NotBefore: jwt.NewNumericDate(time.Now()),
	})
	ss, err := token.SignedString([]byte(u.jwtCfg.RefreshSecretKey))
	if err != nil {
		log.Printf("error sign refresh token. Error: %s", err.Error())
		return "", &authException.SignRefreshToken{}
	}

	return ss, nil
}
