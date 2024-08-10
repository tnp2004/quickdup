package notesException

type InsertNoteLogin struct{}

func (e *InsertNoteLogin) Error() string {
	return "add note failed"
}

type InsertNoteNoLogin struct{}

func (e *InsertNoteNoLogin) Error() string {
	return "add note failed"
}
