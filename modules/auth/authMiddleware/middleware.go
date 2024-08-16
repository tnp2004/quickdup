package authMiddleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/configs"
	"github.com/tnp2004/quickdup/modules/auth/authException"
	"github.com/tnp2004/quickdup/pkg/utils"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cfg := configs.NewConfig()
		accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
		token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, &authException.UnexpectedSigningMethod{}
			}
			return []byte(cfg.Auth.Jwt.AccessSecretKey), nil
		})
		if err != nil {
			log.Printf("error authorization. Error: %s", err.Error())
			return utils.MessageResp(c, http.StatusUnauthorized, "unauthorized")
		}

		if !token.Valid {
			log.Println("error invalid token")
			return utils.MessageResp(c, http.StatusUnauthorized, "unauthorized")
		}

		return next(c)
	}
}
