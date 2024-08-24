package notesException

import "fmt"

type GetNoteBody struct {
	Code string
}

func (e *GetNoteBody) Error() string {
	return fmt.Sprintf("code %s not found", e.Code)
}
