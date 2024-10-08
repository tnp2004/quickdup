package servers

import (
	"github.com/tnp2004/quickdup/modules/notes/notesController"
	"github.com/tnp2004/quickdup/modules/notes/notesRepository"
	"github.com/tnp2004/quickdup/modules/notes/notesUsecase"
)

func (s *Server) registerNotesRouter() {
	r := s.server.Group("/api/v1/notes")

	notesRepository := notesRepository.NewNotesRepository(s.db)
	notesUsecase := notesUsecase.NewNotesUsecase(notesRepository)
	notesController := notesController.NewNotesController(notesUsecase)

	r.POST("/", notesController.AddNewNote)
	r.POST("/code", notesController.GenerateCode)
	r.GET("/:code", notesController.SearchNoteBody)
}
