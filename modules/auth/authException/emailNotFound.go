package authException

import "fmt"

type EmailNotFound struct {
	Email string
}

func (e *EmailNotFound) Error() string {
	return fmt.Sprintf("user email %s not found", e.Email)
}
