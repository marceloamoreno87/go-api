package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int32  `json:"role_id"`
}

type CreateUserOutputDTO struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    int32     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewCreateUserUseCase(userRepository repository.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInputDTO) (output CreateUserOutputDTO, err error) {

	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleId)
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	u, err := uc.UserRepository.CreateUser(user)
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	output = CreateUserOutputDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		RoleId:    u.RoleId,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return
}
