package entities

import "time"

type NoteEntity struct {
	ID        string
	UserID    string
	Blocks    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
