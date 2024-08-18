package notesUsecase

import (
	"github.com/tnp2004/quickdup/modules/models"
	"github.com/tnp2004/quickdup/modules/notes/notesException"
	"github.com/tnp2004/quickdup/modules/notes/notesRepository"
)

type NotesUsecase interface {
	AddNewNote(req *models.InsertNoteRequest) (*models.NoteCode, error)
	GenerateCode(req *models.NoteCode) (*models.NoteCode, error)
}

type notesUsecaseImpl struct {
	notesRepository notesRepository.NotesRepository
}

func NewNotesUsecase(notesRepository notesRepository.NotesRepository) NotesUsecase {
	return &notesUsecaseImpl{notesRepository}
}

func (u *notesUsecaseImpl) AddNewNote(req *models.InsertNoteRequest) (*models.NoteCode, error) {
	resp := new(models.NoteDataID)
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
	genCodeReq := &models.NoteCode{
		NoteID: resp.NoteID,
	}
	codeResp, err := u.GenerateCode(genCodeReq)
	if err != nil {
		return nil, err
	}

	return codeResp, nil
}

func (u *notesUsecaseImpl) GenerateCode(req *models.NoteCode) (*models.NoteCode, error) {
	id, err := u.notesRepository.InsertNoteCode(req)
	if err != nil {
		return nil, &notesException.GenerateCode{}
	}

	return id, nil
}
