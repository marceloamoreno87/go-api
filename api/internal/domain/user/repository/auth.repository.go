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
		DB: config.NewSqlc(config.DB),
	}
}

func (repo *AuthRepository) CreateAuth(auth entityInterface.AuthInterface) (err error) {
	return repo.DB.GetDbQueries().CreateAuth(context.Background(), db.CreateAuthParams{
		UserID:                auth.GetUserID(),
		Token:                 auth.GetToken(),
		RefreshToken:          auth.GetRefreshToken(),
		TokenExpiresIn:        auth.GetTokenExpiresIn(),
		RefreshTokenExpiresIn: auth.GetRefreshTokenExpiresIn(),
	})
}

func (repo *AuthRepository) GetAuthByUserID(userId int32) (output entityInterface.AuthInterface, err error) {
	a, err := repo.DB.GetDbQueries().GetAuthByUserID(context.Background(), userId)
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:                a.UserID,
		Token:                 a.Token,
		RefreshToken:          a.RefreshToken,
		TokenExpiresIn:        a.TokenExpiresIn,
		RefreshTokenExpiresIn: a.RefreshTokenExpiresIn,
		Active:                a.Active,
		CreatedAt:             a.CreatedAt,
		UpdatedAt:             a.UpdatedAt,
	}
	return
}

func (repo *AuthRepository) UpdateAuthRevokeByUserID(userId int32) (err error) {
	return repo.DB.GetDbQueries().UpdateAuthRevokeByUserID(context.Background(), userId)
}

func (repo *AuthRepository) GetAuthByToken(userId int32, token string) (output entityInterface.AuthInterface, err error) {
	a, err := repo.DB.GetDbQueries().GetAuthByToken(context.Background(), db.GetAuthByTokenParams{
		Token:  token,
		UserID: userId,
	})
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:                a.UserID,
		Token:                 a.Token,
		RefreshToken:          a.RefreshToken,
		TokenExpiresIn:        a.TokenExpiresIn,
		RefreshTokenExpiresIn: a.RefreshTokenExpiresIn,
		Active:                a.Active,
		CreatedAt:             a.CreatedAt,
		UpdatedAt:             a.UpdatedAt,
	}
	return
}

func (repo *AuthRepository) GetAuthByRefreshToken(userId int32, refreshToken string) (output entityInterface.AuthInterface, err error) {
	a, err := repo.DB.GetDbQueries().GetAuthByRefreshToken(context.Background(), db.GetAuthByRefreshTokenParams{
		RefreshToken: refreshToken,
		UserID:       userId,
	})
	if err != nil {
		return
	}
	output = &entity.Auth{
		UserID:                a.UserID,
		Token:                 a.Token,
		RefreshToken:          a.RefreshToken,
		TokenExpiresIn:        a.TokenExpiresIn,
		RefreshTokenExpiresIn: a.RefreshTokenExpiresIn,
		Active:                a.Active,
		CreatedAt:             a.CreatedAt,
		UpdatedAt:             a.UpdatedAt,
	}
	return
}
