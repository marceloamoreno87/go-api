package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRoleInputDTO struct {
	ID int32 `json:"id"`
}

type GetRoleOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetRoleUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewGetRoleUseCase() *GetRoleUseCase {
	return &GetRoleUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *GetRoleUseCase) Execute(input GetRoleInputDTO) (output GetRoleOutputDTO, err error) {

	role, err := uc.repo.GetRole(input.ID)
	if err != nil {
		return
	}

	output = GetRoleOutputDTO{
		ID:           role.GetID(),
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
		CreatedAt:    role.GetCreatedAt(),
		UpdatedAt:    role.GetUpdatedAt(),
	}

	return
}
