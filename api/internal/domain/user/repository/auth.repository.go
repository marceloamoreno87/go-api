package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type AuthRepository struct {
	config.SQLCInterface
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (repo *AuthRepository) CreateToken(auth entityInterface.AuthInterface) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateToken(context.Background(), db.CreateTokenParams{
		UserID:       auth.GetUserID(),
		Token:        auth.GetToken(),
		RefreshToken: auth.GetRefreshToken(),
		ExpiresIn:    auth.GetExpiresIn(),
	})
	return
}

func (repo *AuthRepository) GetTokenByUser() (auth entityInterface.AuthInterface, err error) {
	t, err := repo.GetDbQueries().GetTokenByUser(context.Background(), auth.GetUserID())
	if err != nil {
		return nil, err
	}
	auth = &entity.Auth{
		UserID:       t.UserID,
		Token:        t.Token,
		RefreshToken: t.RefreshToken,
		ExpiresIn:    t.ExpiresIn,
	}
	return
}

func (repo *AuthRepository) RevokeTokenByUser(auth entityInterface.AuthInterface) (err error) {
	return repo.GetDbQueries().WithTx(repo.GetTx()).RevokeTokenByUser(context.Background(), auth.GetUserID())
}
