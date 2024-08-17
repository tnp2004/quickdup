package authRepository

import (
	"github.com/tnp2004/quickdup/modules/auth/authModels"
	"github.com/tnp2004/quickdup/pkg/databases"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type AuthRository interface {
	QueryLoginData(req *authModels.LoginRequest) (*authModels.Authentication, error)
	InsertAuthorizationCredentials(req *authModels.AuthorizationCredentials) error
}

type authRepositoryImpl struct {
	db databases.Database
}

func NewAuthRepository(db databases.Database) AuthRository {
	return &authRepositoryImpl{db}
}

func (a *authRepositoryImpl) QueryLoginData(req *authModels.LoginRequest) (*authModels.Authentication, error) {
	query := "SELECT id,password FROM users WHERE email = $1;"
	resp := new(authModels.Authentication)
	args := utils.MakeArgs(req.Email)

	if err := a.db.QueryRowTransaction(query, args, &resp.UserID, &resp.HashPassword); err != nil {
		return nil, err
	}
	resp.Password = req.Password
	return resp, nil
}

func (a *authRepositoryImpl) InsertAuthorizationCredentials(req *authModels.AuthorizationCredentials) error {
	query := "INSERT INTO auth (access_token,refresh_token) VALUES ($1,$2);"
	args := utils.MakeArgs(req.AccessToken, req.RefreshToken)

	if err := a.db.ExecTransaction(query, args); err != nil {
		return err
	}

	return nil
}
