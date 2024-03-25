package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"golang.org/x/net/context"
)

type DeleteRolePermissionByRoleIDInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type DeleteRolePermissionByRoleIDOutputDTO struct {
	RoleID int32 `json:"role_id"`
}

type DeleteRolePermissionByRoleIDUseCase struct {
	repo repository.RolePermissionrepository
}

func NewDeleteRolePermissionByRoleIDUseCase(db config.SQLCInterface) *DeleteRolePermissionByRoleIDUseCase {
	return &DeleteRolePermissionByRoleIDUseCase{
		repo: repository.NewRolePermissionRepository(db),
	}
}

func (uc *DeleteRolePermissionByRoleIDUseCase) Execute(ctx context.Context, input DeleteRolePermissionByRoleIDInputDTO) (output DeleteRolePermissionByRoleIDOutputDTO, err error) {
	err = uc.repo.DeleteRolePermissionByRoleID(ctx, input.RoleID)
	if err != nil {
		return
	}

	output = DeleteRolePermissionByRoleIDOutputDTO{
		RoleID: input.RoleID,
	}

	return
}
