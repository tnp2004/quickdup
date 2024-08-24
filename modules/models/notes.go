package models

type InsertNoteRequest struct {
	UserID string `json:"userID" validate:"omitempty"`
	Blocks string `json:"blocks" validate:"required"`
}

type NoteDataID struct {
	NoteID string `json:"noteID"`
	UserID string `json:"userID"`
}

type NoteCode struct {
	NoteID string `json:"noteID" validate:"required"`
	Code   string `json:"code"`
}

type NoteBlocks struct {
	Blocks string `json:"blocks"`
}
