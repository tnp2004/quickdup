package usersRepository

import (
	"log"

	"github.com/tnp2004/quickdup/modules/models"
	usersexception "github.com/tnp2004/quickdup/modules/users/usersException"
	"github.com/tnp2004/quickdup/pkg/databases"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type UsersRepository interface {
	InsertUser(req *models.UserRegisterRequest) error
}

type usersRepositoryImpl struct {
	db databases.Database
}

func NewUsersRepository(db databases.Database) UsersRepository {
	return &usersRepositoryImpl{db}
}

func (r *usersRepositoryImpl) InsertUser(req *models.UserRegisterRequest) error {
	args := utils.MakeArgs(req.Username, req.Email, req.Password)
	if err := r.db.ExecTransaction("INSERT INTO users (username,email,password) VALUES ($1,$2,$3);", args); err != nil {
		log.Printf("error insert user email %s. Error: %s", req.Email, err.Error())
		return &usersexception.InsertUser{Email: req.Email}
	}

	return nil
}
