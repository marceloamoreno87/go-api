package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
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
	repo repository.RoleRepositoryInterface
}

func NewCreateRoleUseCase(repo repository.RoleRepositoryInterface) *CreateRoleUseCase {
	return &CreateRoleUseCase{
		repo: repo,
	}
}

func (uc *CreateRoleUseCase) Execute(input CreateRoleInputDTO) (err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	if err = uc.repo.CreateRole(role); err != nil {
		return
	}

	return
}
