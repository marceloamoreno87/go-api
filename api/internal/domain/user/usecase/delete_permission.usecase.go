package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeletePermissionInputDTO struct {
	ID int32 `json:"id"`
}

type DeletePermissionOutputDTO struct {
	ID int32 `json:"id"`
}

type DeletePermissionUseCase struct {
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewDeletePermissionUseCase() *DeletePermissionUseCase {
	return &DeletePermissionUseCase{
		repo: repository.NewPermissionRepository(),
	}
}

func (uc *DeletePermissionUseCase) Execute(input DeletePermissionInputDTO) (output DeletePermissionOutputDTO, err error) {
	err = uc.repo.DeletePermission(input.ID)
	output = DeletePermissionOutputDTO{
		input.ID,
	}
	return
}
