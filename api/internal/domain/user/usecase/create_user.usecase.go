package usecase

import (
	"context"

	"github.com/marceloamoreno/izimoney/internal/domain/user/entity"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

type CreateUserUseCase struct {
	UserRepository *db.Queries
}

func NewCreateUserUseCase(userRepository *db.Queries) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(CreateUserParams db.CreateUserParams) (repo db.User, err error) {

	user, err := entity.NewUser(CreateUserParams.Username, CreateUserParams.Password, CreateUserParams.Photo)
	if err != nil {
		return db.User{}, err
	}

	repo, err = uc.UserRepository.CreateUser(context.Background(), db.CreateUserParams{
		Username: user.GetUserName(),
		Password: user.GetPassword(),
		Photo:    user.GetPhoto(),
	})
	if err != nil {
		return db.User{}, err
	}
	return
}
