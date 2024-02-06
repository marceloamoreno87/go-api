package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type DeleteRoleInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteRoleOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DeleteRoleUseCase struct {
	RoleRepository repository.RoleRepositoryInterface
}

func NewDeleteRoleUseCase(roleRepository repository.RoleRepositoryInterface) *DeleteRoleUseCase {
	return &DeleteRoleUseCase{
		RoleRepository: roleRepository,
	}
}

func (uc *DeleteRoleUseCase) Execute(input DeleteRoleInputDTO) (output DeleteRoleOutputDTO, err error) {

	role, err := uc.RoleRepository.GetRole(input.ID)
	if err != nil {
		return DeleteRoleOutputDTO{}, err
	}

	u, err := uc.RoleRepository.DeleteRole(role.GetID())
	if err != nil {
		return DeleteRoleOutputDTO{}, err
	}

	output = DeleteRoleOutputDTO{
		Name:         u.Name,
		InternalName: u.InternalName,
		Description:  u.Description,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
	return
}
