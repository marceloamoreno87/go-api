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
	RoleRepository repository.RoleRepositoryInterface
}

func NewCreateRoleUseCase(roleRepository repository.RoleRepositoryInterface) *CreateRoleUseCase {
	return &CreateRoleUseCase{
		RoleRepository: roleRepository,
	}
}

func (uc *CreateRoleUseCase) Execute(input CreateRoleInputDTO) (err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	err = uc.RoleRepository.CreateRole(role)
	if err != nil {
		return
	}

	return
}
