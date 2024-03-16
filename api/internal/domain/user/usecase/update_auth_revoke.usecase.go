package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateAuthRevokeInputDTO struct {
	UserID int32 `json:"user_id"`
}

type UpdateAuthRevokeOutputDTO struct {
	UserID int32 `json:"user_id"`
}

type UpdateAuthRevokeUseCase struct {
	repo repositoryInterface.AuthRepositoryInterface
}

func NewUpdateAuthRevokeUseCase() *UpdateAuthRevokeUseCase {
	return &UpdateAuthRevokeUseCase{
		repo: repository.NewAuthRepository(),
	}
}

func (uc *UpdateAuthRevokeUseCase) Execute(input UpdateAuthRevokeInputDTO) (output UpdateAuthRevokeOutputDTO, err error) {
	err = uc.repo.UpdateAuthRevokeByUserID(input.UserID)
	if err != nil {
		return
	}

	output = UpdateAuthRevokeOutputDTO{
		UserID: input.UserID,
	}

	return
}
