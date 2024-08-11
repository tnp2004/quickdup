package servers

import (
	"github.com/tnp2004/quickdup/modules/users/usersController"
	"github.com/tnp2004/quickdup/modules/users/usersRepository"
	"github.com/tnp2004/quickdup/modules/users/usersUsecase"
)

func (s *Server) registerUsersRouter() {
	r := s.server.Group("/api/v1/users")

	repository := usersRepository.NewUsersRepository(s.db)
	usecase := usersUsecase.NewUsersUsecase(repository)
	controller := usersController.NewUsersController(usecase)

	r.POST("", controller.AddNewUser)
}
