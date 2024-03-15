package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateRoleInputDTO struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type CreateRoleOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateRoleUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewCreateRoleUseCase() *CreateRoleUseCase {
	return &CreateRoleUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *CreateRoleUseCase) Execute(input CreateRoleInputDTO) (output CreateRoleOutputDTO, err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	r, err := uc.repo.CreateRole(role)
	if err != nil {
		return
	}
	output = CreateRoleOutputDTO{
		ID:           r.GetID(),
		Name:         r.GetName(),
		InternalName: r.GetInternalName(),
		Description:  r.GetDescription(),
		CreatedAt:    r.GetCreatedAt(),
		UpdatedAt:    r.GetUpdatedAt(),
	}
	return
}
