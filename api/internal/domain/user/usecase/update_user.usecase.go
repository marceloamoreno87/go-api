package usecase

import (
	"context"

	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

type UpdateUserUseCase struct {
	UserRepository *db.Queries
}

func NewUpdateUserUseCase(userRepository *db.Queries) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UpdateUserUseCase) Execute(UpdateUserParams db.UpdateUserParams) (repo db.User, err error) {
	repo, err = uc.UserRepository.UpdateUser(context.Background(), UpdateUserParams)
	if err != nil {
		return db.User{}, err
	}
	return
}
