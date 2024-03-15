package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
)

type LoginUserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUserOutputDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUserUseCase struct {
}

func NewLoginUserUseCase() *LoginUserUseCase {
	return &LoginUserUseCase{}
}

func (uc *LoginUserUseCase) Execute(input LoginUserInputDTO) (output LoginUserOutputDTO, err error) {

	// new user entity
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}
	if !user.ComparePassword(input.RequestPassword) {
		return
	}
	// generate token
	user.GenerateToken()
	// save token in DB

	// return token

	return
}
