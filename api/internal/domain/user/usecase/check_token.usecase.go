package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
)

type CheckTokenInputDTO struct {
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

type CheckTokenOutputDTO struct {
	Valid bool `json:"valid"`
}

type CheckTokenUseCase struct {
}

func NewCheckTokenUseCase() *CheckTokenUseCase {
	return &CheckTokenUseCase{}
}

func (uc *CheckTokenUseCase) Execute(input CheckTokenInputDTO) (output CheckTokenOutputDTO, err error) {
	auth, err := entity.NewAuth(input.UserID)
	if err != nil {
		return
	}
	auth.SetToken(input.Token)
	output = CheckTokenOutputDTO{Valid: auth.IsValidToken()}
	return
}
