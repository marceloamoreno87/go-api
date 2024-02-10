package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteUserInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteUserOutputDTO struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    int32     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewDeleteUserUseCase(userRepository repository.UserRepositoryInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Execute(input DeleteUserInputDTO) (output DeleteUserOutputDTO, err error) {
	user, err := uc.UserRepository.GetUser(input.ID)
	if err != nil {
		return DeleteUserOutputDTO{}, err
	}

	u, err := uc.UserRepository.DeleteUser(user.GetID())
	if err != nil {
		return DeleteUserOutputDTO{}, err
	}

	output = DeleteUserOutputDTO{
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		RoleId:    u.RoleId,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return
}
