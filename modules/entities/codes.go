package entities

import "time"

type CodeEntity struct {
	ID        string
	NoteID    string
	ExpiredAt time.Time
	CreatedAt time.Time
}
