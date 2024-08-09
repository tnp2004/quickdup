package notesUsecase

import "github.com/tnp2004/quickdup/modules/notes/notesRepository"

type NotesUsecase interface{}

type notesUsecaseImpl struct {
	NotesRepository *notesRepository.NotesRepository
}

func NewNotesUsecase(NotesRepository *notesRepository.NotesRepository) NotesUsecase {
	return &notesUsecaseImpl{NotesRepository}
}
