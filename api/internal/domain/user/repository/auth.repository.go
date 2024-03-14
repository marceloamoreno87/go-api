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

func (repo *AuthRepository) CreateToken(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error) {
	a, err := repo.GetDbQueries().WithTx(repo.GetTx()).CreateToken(context.Background(), db.CreateTokenParams{
		UserID:       auth.GetUserID(),
		Token:        auth.GetToken(),
		RefreshToken: auth.GetRefreshToken(),
		ExpiresIn:    auth.GetExpiresIn(),
	})
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:       a.UserID,
		Token:        a.Token,
		RefreshToken: a.RefreshToken,
		ExpiresIn:    a.ExpiresIn,
		Active:       a.Active,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
	return
}

func (repo *AuthRepository) GetTokenByUser(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error) {
	a, err := repo.GetDbQueries().GetTokenByUser(context.Background(), auth.GetUserID())
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:       a.UserID,
		Token:        a.Token,
		RefreshToken: a.RefreshToken,
		ExpiresIn:    a.ExpiresIn,
		Active:       a.Active,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
	return
}

func (repo *AuthRepository) RevokeTokenByUser(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error) {
	a, err := repo.GetDbQueries().WithTx(repo.GetTx()).RevokeTokenByUser(context.Background(), auth.GetUserID())
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:       a.UserID,
		Token:        a.Token,
		RefreshToken: a.RefreshToken,
		ExpiresIn:    a.ExpiresIn,
		Active:       a.Active,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
	return
}
