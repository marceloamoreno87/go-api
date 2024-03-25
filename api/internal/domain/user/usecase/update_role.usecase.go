package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateRoleInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdateRoleOutputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdateRoleUseCase struct {
	repo repository.Rolerepository
}

func NewUpdateRoleUseCase(db config.SQLCInterface) *UpdateRoleUseCase {
	return &UpdateRoleUseCase{
		repo: repository.NewRoleRepository(db),
	}
}

func (uc *UpdateRoleUseCase) Execute(ctx context.Context, input UpdateRoleInputDTO) (output UpdateRoleOutputDTO, err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}
	err = uc.repo.UpdateRole(ctx, role, input.ID)

	output = UpdateRoleOutputDTO{
		ID:           role.GetID(),
		Name:         role.GetName(),
		InternalName: role.GetInternalName(),
		Description:  role.GetDescription(),
	}
	return
}
