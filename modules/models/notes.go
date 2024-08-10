package models

type InsertNoteRequest struct {
	UserID string `json:"userID" validate:"omitempty"`
	Blocks string `json:"blocks" validate:"required"`
}

type InsertNoteResponse struct {
	ID     string `json:"ID"`
	UserID string `json:"userID"`
}
