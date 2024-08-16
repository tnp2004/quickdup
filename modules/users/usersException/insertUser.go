package usersException

import "fmt"

type InsertUser struct {
	Email string
}

func (e *InsertUser) Error() string {
	return fmt.Sprintf("add user email %s failed", e.Email)
}
