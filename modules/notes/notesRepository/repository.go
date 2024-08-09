package notesRepository

import "github.com/tnp2004/quickdup/pkg/databases"

type NotesRepository interface{}

type notesRepositoryImpl struct {
	db *databases.Database
}

func NewNotesRepositories(db *databases.Database) NotesRepository {
	return &notesRepositoryImpl{db}
}
