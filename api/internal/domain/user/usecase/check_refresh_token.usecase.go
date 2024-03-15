package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
)

type CheckRefreshTokenInputDTO struct {
	UserID       int32  `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type CheckRefreshTokenOutputDTO struct {
	Valid bool `json:"valid"`
}

type CheckRefreshTokenUseCase struct {
}

func NewCheckRefreshTokenUseCase() *CheckRefreshTokenUseCase {
	return &CheckRefreshTokenUseCase{}
}

func (uc *CheckRefreshTokenUseCase) Execute(input CheckRefreshTokenInputDTO) (output CheckRefreshTokenOutputDTO, err error) {
	auth, err := entity.NewAuth(input.UserID)
	if err != nil {
		return
	}
	auth.SetRefreshToken(input.RefreshToken)
	output = CheckRefreshTokenOutputDTO{Valid: auth.IsValidRefreshToken()}
	return
}
