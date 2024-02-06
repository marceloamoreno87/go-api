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
	RoleRepository repository.RoleRepositoryInterface
}

func NewGetRolesUseCase(roleRepository repository.RoleRepositoryInterface) *GetRolesUseCase {
	return &GetRolesUseCase{
		RoleRepository: roleRepository,
	}
}

func (uc *GetRolesUseCase) Execute(input GetRolesInputDTO) (output []GetRolesOutputDTO, err error) {

	roles, err := uc.RoleRepository.GetRoles(input.Limit, input.Offset)
	if err != nil {
		return []GetRolesOutputDTO{}, err
	}

	for _, role := range roles {
		output = append(output, GetRolesOutputDTO{
			ID:           role.ID,
			Name:         role.Name,
			InternalName: role.InternalName,
			Description:  role.Description,
		})
	}

	return
}
