package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserInputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type UpdateUserUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewUpdateUserUseCase(repo repository.UserRepositoryInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repo: repo,
	}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInputDTO) (err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}
	user.SetActive(input.Active)
	if err = uc.repo.UpdateUser(user, input.ID); err != nil {
		return
	}

	return
}
