package usecase

import (
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

func NewUpdateUserValidationUsedUseCase() *UpdateUserValidationUsedUseCase {
	return &UpdateUserValidationUsedUseCase{
		repo: repository.NewUserValidationRepository(),
	}
}

func (uc *UpdateUserValidationUsedUseCase) Execute(input UpdateUserValidationUsedInputDTO) (err error) {
	return uc.repo.UpdateUserValidationUsed(input.ID)
}
