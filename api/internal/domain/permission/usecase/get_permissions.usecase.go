package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
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
	repo repository.PermissionRepositoryInterface
}

func NewGetPermissionsUseCase(repo repository.PermissionRepositoryInterface) *GetPermissionsUseCase {
	return &GetPermissionsUseCase{
		repo: repo,
	}
}

func (uc *GetPermissionsUseCase) Execute(input GetPermissionsInputDTO) (output []GetPermissionsOutputDTO, err error) {
	permissions, err := uc.repo.GetPermissions(input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, permission := range permissions {
		output = append(output, GetPermissionsOutputDTO{
			ID:           permission.ID,
			Name:         permission.Name,
			InternalName: permission.InternalName,
			Description:  permission.Description,
			CreatedAt:    permission.CreatedAt,
			UpdatedAt:    permission.UpdatedAt,
		})
	}

	return
}
