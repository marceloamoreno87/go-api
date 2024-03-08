package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/event"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type RegisterInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type RegisterOutputDTO struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	RoleID    int32     `json:"role_id"`
	AvatarID  int32     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewRegisterUseCase(
	repo repository.UserRepositoryInterface,
) *RegisterUseCase {
	return &RegisterUseCase{
		repo: repo,
	}
}

func (uc *RegisterUseCase) Execute(input RegisterInputDTO) (output RegisterOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, 1, 1)
	if err != nil {
		return
	}

	newUser, err := uc.repo.RegisterUser(user)
	if err != nil {
		return
	}

	userValidation, err := entity.NewUserValidation(newUser)
	if err != nil {
		return
	}

	err = uc.repo.CreateValidationUser(userValidation)
	if err != nil {
		return
	}
	
	go event.NewUserVerifyEmailEvent(userValidation).Send()

	output = RegisterOutputDTO{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Active:    newUser.Active,
		RoleID:    newUser.RoleID,
		AvatarID:  newUser.AvatarID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	return
}
