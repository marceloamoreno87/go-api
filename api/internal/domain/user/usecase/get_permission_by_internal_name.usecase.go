package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetPermissionByInternalNameInputDTO struct {
	InternalName string `json:"internal_name"`
}

type GetPermissionByInternalNameOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetPermissionByInternalNameUseCase struct {
	repo repositoryInterface.PermissionRepositoryInterface
}

func NewGetPermissionByInternalNameUseCase(DB config.SQLCInterface) *GetPermissionByInternalNameUseCase {
	return &GetPermissionByInternalNameUseCase{
		repo: repository.NewPermissionRepository(DB),
	}
}

func (uc *GetPermissionByInternalNameUseCase) Execute(input GetPermissionByInternalNameInputDTO) (output GetPermissionByInternalNameOutputDTO, err error) {
	permission, err := uc.repo.GetPermissionByInternalName(input.InternalName)
	if err != nil {
		return
	}

	output = GetPermissionByInternalNameOutputDTO{
		ID:           permission.GetID(),
		Name:         permission.GetName(),
		InternalName: permission.GetInternalName(),
		Description:  permission.GetDescription(),
		CreatedAt:    permission.GetCreatedAt(),
		UpdatedAt:    permission.GetUpdatedAt(),
	}
	return
}
