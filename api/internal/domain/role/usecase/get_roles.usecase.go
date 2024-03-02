package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
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
	repo repository.RoleRepositoryInterface
}

func NewGetRolesUseCase(repo repository.RoleRepositoryInterface) *GetRolesUseCase {
	return &GetRolesUseCase{
		repo: repo,
	}
}

func (uc *GetRolesUseCase) Execute(input GetRolesInputDTO) (output []GetRolesOutputDTO, err error) {

	roles, err := uc.repo.GetRoles(input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, role := range roles {
		output = append(output, GetRolesOutputDTO{
			ID:           role.ID,
			Name:         role.Name,
			InternalName: role.InternalName,
			Description:  role.Description,
			CreatedAt:    role.CreatedAt,
			UpdatedAt:    role.UpdatedAt,
		})
	}

	return
}
