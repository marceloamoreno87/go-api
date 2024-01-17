package usecase

import "github.com/marceloamoreno/izimoney/internal/domain/user/repository"

type GetUserInputDTO struct {
	ID int64 `json:"id"`
}

type GetUserOutputDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
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
		Username: user.Username,
		Password: user.Password,
		Photo:    user.Photo,
	}

	return
}
