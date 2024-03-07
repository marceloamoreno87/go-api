package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

func (repo *UserRepository) GetValidationUser(id int32) (userValidation *entity.UserValidation, err error) {
	newUserValidation, err := repo.GetDbQueries().GetValidationUser(context.Background(), id)
	if err != nil {
		return
	}
	userValidation = &entity.UserValidation{
		ID:        newUserValidation.ID,
		UserID:    newUserValidation.UserID,
		Hash:      newUserValidation.Hash,
		ExpiresIn: newUserValidation.ExpiresIn,
		CreatedAt: newUserValidation.CreatedAt,
	}
	return
}

func (repo *UserRepository) GetValidationUserByHash(hash string) (userValidation *entity.UserValidation, err error) {
	newUserValidation, err := repo.GetDbQueries().GetValidationUserByHash(context.Background(), hash)
	if err != nil {
		return
	}
	userValidation = &entity.UserValidation{
		ID:        newUserValidation.ID,
		UserID:    newUserValidation.UserID,
		Hash:      newUserValidation.Hash,
		ExpiresIn: newUserValidation.ExpiresIn,
		CreatedAt: newUserValidation.CreatedAt,
	}
	return
}

func (repo *UserRepository) CreateValidationUser(userValidation *entity.UserValidation) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateValidationUser(context.Background(), db.CreateValidationUserParams{
		UserID:    userValidation.GetUserID(),
		Hash:      userValidation.GetHash(),
		ExpiresIn: userValidation.GetExpiresIn(),
	})
	return
}

func (repo *UserRepository) UpdateValidationUser(userValidation *entity.UserValidation, id int32) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).UpdateValidationUser(context.Background(), db.UpdateValidationUserParams{
		ID:   id,
		Used: userValidation.Used,
	})
	return
}
