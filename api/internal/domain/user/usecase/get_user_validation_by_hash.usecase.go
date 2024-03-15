package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUserValidationByHashInputDTO struct {
	Hash string `json:"hash"`
}

type GetUserValidationByHashOutputDTO struct {
	ID        int32  `json:"id"`
	UserID    int32  `json:"user_id"`
	Hash      string `json:"hash"`
	ExpiresIn int32  `json:"expires_in"`
	Used      bool   `json:"used"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUserValidationByHashUseCase struct {
	repo repositoryInterface.UserValidationRepositoryInterface
}

func NewGetUserValidationByHashUseCase(DB config.SQLCInterface) *GetUserValidationByHashUseCase {
	return &GetUserValidationByHashUseCase{
		repo: repository.NewUserValidationRepository(DB),
	}
}

func (uc *GetUserValidationByHashUseCase) Execute(input GetUserValidationByHashInputDTO) (output GetUserValidationByHashOutputDTO, err error) {
	userValidation, err := uc.repo.GetValidationUserByHash(input.Hash)
	if err != nil {
		return
	}

	if !userValidation.ValidateHashExpiresIn() {
		return
	}

	return
}
