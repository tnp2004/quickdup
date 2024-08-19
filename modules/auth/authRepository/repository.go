package authRepository

import (
	"github.com/tnp2004/quickdup/modules/auth/authModels"
	"github.com/tnp2004/quickdup/pkg/databases"
	"github.com/tnp2004/quickdup/pkg/utils"
)

type AuthRository interface {
	QueryLoginData(req *authModels.LoginRequest) (*authModels.Authentication, error)
	InsertAuthorizationCredentials(req *authModels.AuthorizationCredentials) error
	DeleteCredential(refreshToken string) error
	IsExistsCredential(accessToken string) error
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

	if err := a.db.QueryRow(query, args, &resp.UserID, &resp.HashPassword); err != nil {
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

func (a *authRepositoryImpl) IsExistsCredential(accessToken string) error {
	query := "SELECT id FROM auth WHERE access_token = $1 LIMIT 1;"
	args := utils.MakeArgs(accessToken)

	id := new(string)
	if err := a.db.QueryRow(query, args, &id); err != nil {
		return err
	}

	return nil
}

func (a *authRepositoryImpl) DeleteCredential(refreshToken string) error {
	query := "DELETE FROM auth WHERE refresh_token = $1;"
	args := utils.MakeArgs(refreshToken)

	if err := a.db.ExecTransaction(query, args); err != nil {
		return err
	}

	return nil
}
