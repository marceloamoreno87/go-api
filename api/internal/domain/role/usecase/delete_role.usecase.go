package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type DeleteRoleInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteRoleUseCase struct {
	RoleRepository repository.RoleRepositoryInterface
}

func NewDeleteRoleUseCase(roleRepository repository.RoleRepositoryInterface) *DeleteRoleUseCase {
	return &DeleteRoleUseCase{
		RoleRepository: roleRepository,
	}
}

func (uc *DeleteRoleUseCase) Execute(input DeleteRoleInputDTO) (err error) {

	role, err := uc.RoleRepository.GetRole(input.ID)
	if err != nil {
		return
	}

	err = uc.RoleRepository.DeleteRole(role.GetID())
	if err != nil {
		return
	}

	return
}
