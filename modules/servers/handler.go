package servers

func (s *Server) registerModuleRouters() {
	s.registerHealthRouter()
	s.registerNotesRouter()
	s.registerUsersRouter()
}
