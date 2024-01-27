package usecase

import "github.com/marceloamoreno/goapi/internal/domain/user/repository"

type GetUserInputDTO struct {
	ID int64 `json:"id"`
}

type GetUserOutputDTO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewGetUserUseCase(userRepository repository.UserRepositoryInterface) *GetUserUseCase {
	return &GetUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetUserUseCase) Execute(input GetUserInputDTO) (output GetUserOutputDTO, err error) {
	user, err := uc.UserRepository.GetUser(input.ID)
	if err != nil {
		return GetUserOutputDTO{}, err
	}

	output = GetUserOutputDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return
}
