package utils

import (
	"log"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func newValidator() *validator.Validate {
	once.Do(func() {
		validatorInstance = validator.New(validator.WithRequiredStructEnabled())
	})

	return validatorInstance
}

func BindRequestBody(c echo.Context, data any) error {
	if err := c.Bind(&data); err != nil {
		log.Printf("error bind request body. Error: %s", err.Error())
		return err
	}

	if err := newValidator().Struct(data); err != nil {
		log.Printf("error validate body. Error: %s", err.Error())
		return err
	}

	return nil
}
