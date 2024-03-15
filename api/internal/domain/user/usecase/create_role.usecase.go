package usecase

import (
	"github.com/marceloamoreno/goapi/config"
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
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type CreateRoleUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewCreateRoleUseCase(DB config.SQLCInterface) *CreateRoleUseCase {
	return &CreateRoleUseCase{
		repo: repository.NewRoleRepository(DB),
	}
}

func (uc *CreateRoleUseCase) Execute(input CreateRoleInputDTO) (output CreateRoleOutputDTO, err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}

	err = uc.repo.CreateRole(role)
	output = CreateRoleOutputDTO{
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	}
	return
}
