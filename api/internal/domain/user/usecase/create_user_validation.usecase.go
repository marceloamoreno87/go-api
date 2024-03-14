package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateUserValidationInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

type CreateUserValidationUseCase struct {
	repo repositoryInterface.UserValidationRepositoryInterface
}

func NewCreateUserValidationUseCase() *CreateUserValidationUseCase {
	return &CreateUserValidationUseCase{
		repo: repository.NewUserValidationRepository(),
	}
}

func (uc *CreateUserValidationUseCase) Execute(input CreateUserValidationInputDTO) (err error) {
	// userValidation, err := entity.NewUserValidation(newUser)
	// if err != nil {
	// 	return
	// }

	// err = uc.repo.CreateValidationUser(userValidation)
	// if err != nil {
	// 	return
	// }

	return
}
