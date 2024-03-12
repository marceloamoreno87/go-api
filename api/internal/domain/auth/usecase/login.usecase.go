package usecase

import (
	"errors"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	Token string `json:"token"`
}

type LoginUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewLoginUseCase(repo repository.UserRepositoryInterface) *LoginUseCase {
	return &LoginUseCase{
		repo: repo,
	}
}

func (uc *LoginUseCase) Execute(input LoginInputDTO) (output LoginOutputDTO, err error) {
	user, err := uc.repo.GetUserByEmail(input.Email)
	if err != nil {
		slog.Info("err", err)
		return
	}

	newUser := &entity.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Active:    user.Active,
		RoleID:    user.RoleID,
		AvatarID:  user.AvatarID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if !newUser.ComparePassword(input.Password) {
		return LoginOutputDTO{}, errors.New("invalid credentials")
	}

	newUser.GenerateToken()

	return LoginOutputDTO{Token: newUser.Token}, nil
}
