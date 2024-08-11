package notesUsecase

import (
	"github.com/tnp2004/quickdup/modules/models"
	"github.com/tnp2004/quickdup/modules/notes/notesRepository"
)

type NotesUsecase interface {
	AddNewNote(req *models.InsertNoteRequest) (*models.InsertNoteResponse, error)
}

type notesUsecaseImpl struct {
	notesRepository notesRepository.NotesRepository
}

func NewNotesUsecase(notesRepository notesRepository.NotesRepository) NotesUsecase {
	return &notesUsecaseImpl{notesRepository}
}

func (u *notesUsecaseImpl) AddNewNote(req *models.InsertNoteRequest) (*models.InsertNoteResponse, error) {
	resp := new(models.InsertNoteResponse)
	var err error

	if req.UserID == "" {
		// no login
		resp, err = u.notesRepository.InsertNoteNoLogin(req)
	} else {
		// login
		resp, err = u.notesRepository.InsertNoteLogin(req)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
