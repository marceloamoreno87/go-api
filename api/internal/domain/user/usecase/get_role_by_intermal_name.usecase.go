package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRoleByInternalNameInputDTO struct {
	InternalName string `json:"internal_name"`
}

type GetRoleByInternalNameOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetRoleByInternalNameUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewGetRoleByInternalNameUseCase() *GetRoleByInternalNameUseCase {
	return &GetRoleByInternalNameUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *GetRoleByInternalNameUseCase) Execute(input GetRoleByInternalNameInputDTO) (output GetRoleByInternalNameOutputDTO, err error) {

	role, err := uc.repo.GetRoleByInternalName(input.InternalName)
	if err != nil {
		return
	}

	output = GetRoleByInternalNameOutputDTO{
		ID:           role.GetID(),
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
		CreatedAt:    role.GetCreatedAt(),
		UpdatedAt:    role.GetUpdatedAt(),
	}

	return
}
