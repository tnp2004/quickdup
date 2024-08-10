package entities

import "time"

type UserEntity struct {
	ID        string
	Username  string
	Email     string
	Passowrd  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
