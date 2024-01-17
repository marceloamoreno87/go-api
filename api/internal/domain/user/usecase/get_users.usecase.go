package usecase

import (
	"github.com/marceloamoreno/izimoney/internal/domain/user/entity"
	"github.com/marceloamoreno/izimoney/internal/domain/user/repository"
)

type GetUsersInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersOutputDTO struct {
	Users []*entity.User
}

type GetUsersUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewGetUsersUseCase(userRepository repository.UserRepositoryInterface) *GetUsersUseCase {
	return &GetUsersUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetUsersUseCase) Execute(input GetUsersInputDTO) (output GetUsersOutputDTO, err error) {
	users, err := uc.UserRepository.GetUsers(input.Limit, input.Offset)
	if err != nil {
		return GetUsersOutputDTO{}, err
	}

	output = GetUsersOutputDTO{
		Users: users,
	}

	return
}
