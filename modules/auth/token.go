package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tnp2004/quickdup/modules/auth/authException"
)

func (a *authImpl) genAccessKey(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "quickdup",
		Subject:   userID,
		Audience:  []string{"user"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(a.cfg.AccessTokenExpireDuration))),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})
	ss, err := token.SignedString([]byte(a.cfg.AccessSecretKey))
	if err != nil {
		log.Printf("error sign access token. Error: %s", err.Error())
		return "", &authException.SignAccessToken{}
	}

	return ss, nil
}

func (a *authImpl) genRefreshKey(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "quickdup",
		Subject:   userID,
		Audience:  []string{"user"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(a.cfg.RefreshTokenExpireDuration))),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})
	ss, err := token.SignedString([]byte(a.cfg.AccessSecretKey))
	if err != nil {
		log.Printf("error sign refresh token. Error: %s", err.Error())
		return "", &authException.SignRefreshToken{}
	}

	return ss, nil
}
