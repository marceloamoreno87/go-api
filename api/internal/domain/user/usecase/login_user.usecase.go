package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
)

type LoginUserInputDTO struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	RoleID          int32  `json:"role_id"`
	AvatarID        int32  `json:"avatar_id"`
	RequestPassword string `json:"request_password"`
}
type LoginUserOutputDTO struct {
	Valid bool `json:"valid"`
}

type LoginUserUseCase struct {
}

func NewLoginUserUseCase() *LoginUserUseCase {
	return &LoginUserUseCase{}
}

func (uc *LoginUserUseCase) Execute(input LoginUserInputDTO) (output LoginUserOutputDTO, err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password, input.RoleID, input.AvatarID)
	if err != nil {
		return
	}
	output.Valid = user.ComparePassword(input.RequestPassword)
	return
}
