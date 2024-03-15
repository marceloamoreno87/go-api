package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type AuthRepository struct {
	DB config.SQLCInterface
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		DB: config.Sqcl,
	}
}

func (repo *AuthRepository) CreateAuth(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error) {
	a, err := repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).CreateAuth(context.Background(), db.CreateAuthParams{
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

func (repo *AuthRepository) GetAuthByUser(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error) {
	a, err := repo.DB.GetDbQueries().GetAuthByUser(context.Background(), auth.GetUserID())
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

func (repo *AuthRepository) RevokeAuthByUser(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error) {
	a, err := repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).RevokeAuthByUser(context.Background(), auth.GetUserID())
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
