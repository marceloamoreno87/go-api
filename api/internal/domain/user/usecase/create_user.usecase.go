package usecase

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutputDTO struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
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

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
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
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return
}
