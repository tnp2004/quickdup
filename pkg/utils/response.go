package utils

import "github.com/labstack/echo/v4"

type messageResponse struct {
	Message string `json:"message"`
}

func MessageResp(c echo.Context, status int, message string) error {
	return c.JSON(status, &messageResponse{Message: message})
}
