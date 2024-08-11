package usersUsecase

import (
	"github.com/tnp2004/quickdup/modules/models"
	"github.com/tnp2004/quickdup/modules/users/usersRepository"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type UsersUsecase interface {
	AddNewUser(req *models.UserRegisterRequest) error
}

type usersUsecaseImpl struct {
	usersRepository usersRepository.UsersRepository
}

func NewUsersUsecase(usersRepository usersRepository.UsersRepository) UsersUsecase {
	return &usersUsecaseImpl{usersRepository}
}

func (u *usersUsecaseImpl) AddNewUser(req *models.UserRegisterRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashedPassword

	if err := u.usersRepository.InsertUser(req); err != nil {
		return err
	}

	return nil
}
