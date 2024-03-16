package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetRolesInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetRolesOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetRolesUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewGetRolesUseCase() *GetRolesUseCase {
	return &GetRolesUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *GetRolesUseCase) Execute(input GetRolesInputDTO) (output []GetRolesOutputDTO, err error) {

	roles, err := uc.repo.GetRoles(input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, role := range roles {
		output = append(output, GetRolesOutputDTO{
			ID:           role.GetID(),
			Name:         role.GetName(),
			InternalName: role.GetInternalName(),
			Description:  role.GetDescription(),
			CreatedAt:    role.GetCreatedAt(),
			UpdatedAt:    role.GetUpdatedAt(),
		})
	}

	return
}
