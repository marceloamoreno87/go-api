package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/marceloamoreno/goapi/config"
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
	repo repository.UserValidationrepository
}

func NewGetUserValidationByHashUseCase(db config.SQLCInterface) *GetUserValidationByHashUseCase {
	return &GetUserValidationByHashUseCase{
		repo: repository.NewUserValidationRepository(db),
	}
}

func (uc *GetUserValidationByHashUseCase) Execute(ctx context.Context, input GetUserValidationByHashInputDTO) (output GetUserValidationByHashOutputDTO, err error) {
	userValidation, err := uc.repo.GetUserValidationByHash(ctx, input.Hash)
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
