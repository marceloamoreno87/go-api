package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetUserValidationByUserIDInputDTO struct {
	UserID int32 `json:"user_id"`
}

type GetUserValidationByUserIDOutputDTO struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Hash      string    `json:"hash"`
	ExpiresIn int32     `json:"expires_in"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserValidationByUserIDUseCase struct {
	repo repositoryInterface.UserValidationRepositoryInterface
}

func NewGetUserValidationByUserIDUseCase(db config.SQLCInterface) *GetUserValidationByUserIDUseCase {
	return &GetUserValidationByUserIDUseCase{
		repo: repository.NewUserValidationRepository(db),
	}
}

func (uc *GetUserValidationByUserIDUseCase) Execute(ctx context.Context, input GetUserValidationByUserIDInputDTO) (output GetUserValidationByUserIDOutputDTO, err error) {
	userValidation, err := uc.repo.GetUserValidationByUserID(ctx, input.UserID)
	if err != nil {
		return
	}

	if userValidation.GetUsed() {
		return output, errors.New("hash already used")
	}

	if !userValidation.ValidateHashExpiresIn() {
		return output, errors.New("hash expired")
	}

	output = GetUserValidationByUserIDOutputDTO{
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
