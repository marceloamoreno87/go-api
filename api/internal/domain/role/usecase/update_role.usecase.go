package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type UpdateRoleInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
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

func (uc *UpdateRoleUseCase) Execute(input UpdateRoleInputDTO) (err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}
	err = uc.RoleRepository.UpdateRole(role, uc.ID)
	if err != nil {
		return
	}

	return
}
