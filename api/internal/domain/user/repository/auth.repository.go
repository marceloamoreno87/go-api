package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type AuthRepository struct {
	config.SQLCInterface
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (repo *AuthRepository) CreateToken(auth entityInterface.AuthInterface) (output db.Auth, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateToken(context.Background(), db.CreateTokenParams{
		UserID:       auth.GetUserID(),
		Token:        auth.GetToken(),
		RefreshToken: auth.GetRefreshToken(),
		ExpiresIn:    auth.GetExpiresIn(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *AuthRepository) GetTokenByUser(auth entityInterface.AuthInterface) (output db.Auth, err error) {
	output, err = repo.GetDbQueries().GetTokenByUser(context.Background(), auth.GetUserID())
	if err != nil {
		return
	}
	return
}

func (repo *AuthRepository) RevokeTokenByUser(auth entityInterface.AuthInterface) (output db.Auth, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).RevokeTokenByUser(context.Background(), auth.GetUserID())
	if err != nil {
		return
	}
	return
}
