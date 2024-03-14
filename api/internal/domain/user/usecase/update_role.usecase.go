package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateRoleInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdateRoleUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewUpdateRoleUseCase() *UpdateRoleUseCase {
	return &UpdateRoleUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *UpdateRoleUseCase) Execute(input UpdateRoleInputDTO) (err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}
	if err = uc.repo.UpdateRole(role, input.ID); err != nil {
		return
	}

	return
}
