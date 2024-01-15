package usecase

import (
	"context"

	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

type GetUserUseCase struct {
	UserRepository *db.Queries
}

func NewGetUserUseCase(userRepository *db.Queries) *GetUserUseCase {
	return &GetUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetUserUseCase) Execute(id int64) (repo db.User, err error) {
	repo, err = uc.UserRepository.GetUser(context.Background(), id)
	if err != nil {
		return db.User{}, err
	}
	return
}
