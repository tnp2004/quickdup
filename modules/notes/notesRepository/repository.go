package notesRepository

import (
	"github.com/tnp2004/quickdup/modules/models"
	"github.com/tnp2004/quickdup/modules/notes/notesException"
	"github.com/tnp2004/quickdup/pkg/databases"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type NotesRepository interface {
	InsertNoteLogin(noteEntity *models.InsertNoteRequest) (*models.InsertNoteResponse, error)
	InsertNoteNoLogin(noteEntity *models.InsertNoteRequest) (*models.InsertNoteResponse, error)
}

type notesRepositoryImpl struct {
	db databases.Database
}

func NewNotesRepository(db databases.Database) NotesRepository {
	return &notesRepositoryImpl{db}
}

// login
func (r *notesRepositoryImpl) InsertNoteLogin(req *models.InsertNoteRequest) (*models.InsertNoteResponse, error) {
	resp := new(models.InsertNoteResponse)
	args := utils.MakeArgs(req.UserID, req.Blocks)
	if err := r.db.QueryRowTransaction("INSERT INTO notes (user_id,blocks) VALUES ($1,$2) RETURNING id,user_id;",
		args, &resp.ID, &resp.UserID); err != nil {
		return nil, &notesException.InsertNoteLogin{}
	}

	return resp, nil
}

// no login
func (r *notesRepositoryImpl) InsertNoteNoLogin(req *models.InsertNoteRequest) (*models.InsertNoteResponse, error) {
	resp := new(models.InsertNoteResponse)

	query := "INSERT INTO notes (blocks) VALUES ($1) RETURNING id;"
	args := utils.MakeArgs(req.Blocks)
	if err := r.db.QueryRowTransaction(query, args, &resp.ID); err != nil {
		return nil, &notesException.InsertNoteNoLogin{}
	}

	return resp, nil
}
