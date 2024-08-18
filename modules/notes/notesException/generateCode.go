package notesException

type GenerateCode struct{}

func (e *GenerateCode) Error() string {
	return "generate code failed"
}
