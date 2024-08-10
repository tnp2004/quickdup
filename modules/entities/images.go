package entities

import "time"

type ImageEntity struct {
	ID        string
	NoteID    string
	URL       string
	CreatedAt time.Time
}
