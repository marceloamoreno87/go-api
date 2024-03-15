package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserValidationUsedInputDTO struct {
	ID int32 `json:"user_id"`
}

type UpdateUserValidationUsedOutputDTO struct {
	ID int32 `json:"user_id"`
}

type UpdateUserValidationUsedUseCase struct {
	repo repositoryInterface.UserValidationRepositoryInterface
}

func NewUpdateUserValidationUsedUseCase(DB config.SQLCInterface) *UpdateUserValidationUsedUseCase {
	return &UpdateUserValidationUsedUseCase{
		repo: repository.NewUserValidationRepository(DB),
	}
}

func (uc *UpdateUserValidationUsedUseCase) Execute(input UpdateUserValidationUsedInputDTO) (err error) {
	return uc.repo.UpdateUserValidationUsed(input.ID)
}
