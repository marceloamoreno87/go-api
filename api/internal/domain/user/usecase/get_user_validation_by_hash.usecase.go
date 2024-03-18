package usecase

import (
	"errors"
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUserValidationByHashInputDTO struct {
	Hash string `json:"hash"`
}

type GetUserValidationByHashOutputDTO struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Hash      string    `json:"hash"`
	ExpiresIn int32     `json:"expires_in"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserValidationByHashUseCase struct {
	repo repositoryInterface.UserValidationRepositoryInterface
}

func NewGetUserValidationByHashUseCase() *GetUserValidationByHashUseCase {
	return &GetUserValidationByHashUseCase{
		repo: repository.NewUserValidationRepository(),
	}
}

func (uc *GetUserValidationByHashUseCase) Execute(input GetUserValidationByHashInputDTO) (output GetUserValidationByHashOutputDTO, err error) {
	userValidation, err := uc.repo.GetUserValidationByHash(input.Hash)
	if err != nil {
		return
	}

	if userValidation.GetUsed() {
		return output, errors.New("hash already used")
	}

	if !userValidation.ValidateHashExpiresIn() {
		return output, errors.New("hash expired")
	}

	output = GetUserValidationByHashOutputDTO{
		ID:        userValidation.GetID(),
		UserID:    userValidation.GetUserID(),
		Hash:      userValidation.GetHash(),
		ExpiresIn: userValidation.GetExpiresIn(),
		Used:      userValidation.GetUsed(),
		CreatedAt: userValidation.GetCreatedAt(),
		UpdatedAt: userValidation.GetUpdatedAt(),
	}
	return
}
