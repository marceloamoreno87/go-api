package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetPermissionsInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetPermissionsOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetPermissionsUseCase struct {
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewGetPermissionsUseCase() *GetPermissionsUseCase {
	return &GetPermissionsUseCase{
		repo: repository.NewPermissionRepository(),
	}
}

func (uc *GetPermissionsUseCase) Execute(input GetPermissionsInputDTO) (output []GetPermissionsOutputDTO, err error) {
	permissions, err := uc.repo.GetPermissions(input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, permission := range permissions {
		output = append(output, GetPermissionsOutputDTO{
			ID:           permission.GetID(),
			Name:         permission.GetName(),
			InternalName: permission.GetInternalName(),
			Description:  permission.GetDescription(),
			CreatedAt:    permission.GetCreatedAt(),
			UpdatedAt:    permission.GetUpdatedAt(),
		})
	}

	return
}
