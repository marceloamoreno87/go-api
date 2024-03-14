package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type UserValidationRepository struct {
	config.SQLCInterface
}

func NewUserValidationRepository() *UserValidationRepository {
	return &UserValidationRepository{}
}

func (repo *UserValidationRepository) GetValidationUser(id int32) (output db.UsersValidation, err error) {
	output, err = repo.GetDbQueries().GetValidationUser(context.Background(), id)
	if err != nil {
		return
	}
	return
}

func (repo *UserValidationRepository) GetValidationUserByHash(hash string) (output db.UsersValidation, err error) {
	output, err = repo.GetDbQueries().GetValidationUserByHash(context.Background(), hash)
	if err != nil {
		return
	}
	return
}

func (repo *UserValidationRepository) CreateValidationUser(userValidation entityInterface.UserValidationInterface) (output db.UsersValidation, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateValidationUser(context.Background(), db.CreateValidationUserParams{
		UserID:    userValidation.GetUserID(),
		Hash:      userValidation.GetHash(),
		ExpiresIn: userValidation.GetExpiresIn(),
	})
	if err != nil {
		return
	}
	return
}

func (repo *UserValidationRepository) UpdateUserValidationUsed(id int32) (output db.UsersValidation, err error) {
	output, err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateUserValidationUsed(context.Background(), id)
	if err != nil {
		return
	}
	return
}
