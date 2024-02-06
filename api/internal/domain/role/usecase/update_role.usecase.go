package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type UpdateRoleInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdateRoleOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateRoleUseCase struct {
	RoleRepository repository.RoleRepositoryInterface
	ID             int32
}

func NewUpdateRoleUseCase(roleRepository repository.RoleRepositoryInterface, id int32) *UpdateRoleUseCase {
	return &UpdateRoleUseCase{
		RoleRepository: roleRepository,
		ID:             id,
	}
}

func (uc *UpdateRoleUseCase) Execute(input UpdateRoleInputDTO) (output UpdateRoleOutputDTO, err error) {

	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return UpdateRoleOutputDTO{}, err
	}

	role.SetID(input.ID)

	u, err := uc.RoleRepository.UpdateRole(role, uc.ID)
	if err != nil {
		return UpdateRoleOutputDTO{}, err
	}

	output = UpdateRoleOutputDTO{
		ID:           u.ID,
		Name:         u.Name,
		InternalName: u.InternalName,
		Description:  u.Description,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}

	return
}
