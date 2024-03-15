package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserPasswordInputDTO struct {
	ID       int32  `json:"id"`
	Password string `json:"password"`
}

type UpdateUserPasswordOutputDTO struct {
	ID       int32  `json:"id"`
	Password string `json:"password"`
}

type UpdateUserPasswordUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdateUserPasswordUseCase(DB config.SQLCInterface) *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		repo: repository.NewUserRepository(DB),
	}
}

func (uc *UpdateUserPasswordUseCase) Execute(input UpdateUserPasswordInputDTO) (output UpdateUserPasswordOutputDTO, err error) {
	err = uc.repo.UpdateUserPassword(input.ID, input.Password)
	output = UpdateUserPasswordOutputDTO{
		ID:       input.ID,
		Password: input.Password,
	}

	return
}
