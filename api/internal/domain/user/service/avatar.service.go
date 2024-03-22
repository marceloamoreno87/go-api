package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarService struct {
	db                  config.SQLCInterface
	GetAvatarUseCase    usecaseInterface.GetAvatarUseCaseInterface
	GetAvatarsUseCase   usecaseInterface.GetAvatarsUseCaseInterface
	CreateAvatarUseCase usecaseInterface.CreateAvatarUseCaseInterface
	UpdateAvatarUseCase usecaseInterface.UpdateAvatarUseCaseInterface
	DeleteAvatarUseCase usecaseInterface.DeleteAvatarUseCaseInterface
}

func NewAvatarService() *AvatarService {
	db := config.NewSqlc(config.DB)
	return &AvatarService{
		db:                  db,
		GetAvatarUseCase:    usecase.NewGetAvatarUseCase(db),
		GetAvatarsUseCase:   usecase.NewGetAvatarsUseCase(db),
		CreateAvatarUseCase: usecase.NewCreateAvatarUseCase(db),
		UpdateAvatarUseCase: usecase.NewUpdateAvatarUseCase(db),
		DeleteAvatarUseCase: usecase.NewDeleteAvatarUseCase(db),
	}
}

func (s *AvatarService) GetAvatar(ctx context.Context, input request.RequestGetAvatar) (output usecase.GetAvatarOutputDTO, err error) {
	output, err = s.GetAvatarUseCase.Execute(ctx, usecase.GetAvatarInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar found")
	return
}

func (s *AvatarService) GetAvatars(ctx context.Context, input request.RequestGetAvatars) (output []usecase.GetAvatarsOutputDTO, err error) {
	output, err = s.GetAvatarsUseCase.Execute(ctx, usecase.GetAvatarsInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatars found")
	return
}

func (s *AvatarService) CreateAvatar(ctx context.Context, input request.RequestCreateAvatar) (output usecase.CreateAvatarOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.CreateAvatarUseCase.Execute(ctx, usecase.CreateAvatarInputDTO{SVG: input.SVG})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(ctx context.Context, input request.RequestUpdateAvatar) (output usecase.UpdateAvatarOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.UpdateAvatarUseCase.Execute(ctx, usecase.UpdateAvatarInputDTO{ID: input.ID, SVG: input.SVG})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(ctx context.Context, input request.RequestDeleteAvatar) (output usecase.DeleteAvatarOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.DeleteAvatarUseCase.Execute(ctx, usecase.DeleteAvatarInputDTO{ID: input.ID})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	slog.Info("Avatar deleted")
	return
}
