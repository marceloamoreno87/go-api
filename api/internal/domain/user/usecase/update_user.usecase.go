package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserInputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type UpdateUserOutputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type UpdateUserUseCase struct {
	repo repositoryInterface.UserRepositoryInterface
}

func NewUpdateUserUseCase(db config.SQLCInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repo: repository.NewUserRepository(db),
	}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, input UpdateUserInputDTO) (output UpdateUserOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return
	}
	user.HashPassword()
	user.SetAvatarID(input.AvatarID)
	user.SetRoleID(input.RoleID)
	err = uc.repo.UpdateUser(ctx, user, input.ID)

	output = UpdateUserOutputDTO{
		ID:       user.GetID(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Active:   user.GetActive(),
		RoleID:   user.GetRoleID(),
		AvatarID: user.GetAvatarID(),
	}
	return
}
