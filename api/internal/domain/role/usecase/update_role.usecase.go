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
	repo repository.RoleRepositoryInterface
}

func NewUpdateRoleUseCase(repo repository.RoleRepositoryInterface) *UpdateRoleUseCase {
	return &UpdateRoleUseCase{
		repo: repo,
	}
}

func (uc *UpdateRoleUseCase) Execute(input UpdateRoleInputDTO) (err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}
	if err = uc.repo.UpdateRole(role, role.ID); err != nil {
		return
	}

	return
}
