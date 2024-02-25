package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type UpdateUserUseCase struct {
	repo repository.UserRepositoryInterface
	ID   int32
}

func NewUpdateUserUseCase(repo repository.UserRepositoryInterface, id int32) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repo: repo,
		ID:   id,
	}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInputDTO) (err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}

	err = uc.repo.UpdateUser(user, uc.ID)
	if err != nil {
		return
	}

	return
}
