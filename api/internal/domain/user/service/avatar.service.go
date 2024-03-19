package service

import (
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RequestCreateAvatarInputDTO struct {
	SVG string `json:"svg"`
}

type RequestUpdateAvatarInputDTO struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type RequestGetAvatarInputDTO struct {
	ID int32 `json:"id"`
}

type RequestGetAvatarsInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type RequestDeleteAvatarInputDTO struct {
	ID int32 `json:"id"`
}

type AvatarService struct {
	GetAvatarUseCase    usecaseInterface.GetAvatarUseCaseInterface
	GetAvatarsUseCase   usecaseInterface.GetAvatarsUseCaseInterface
	CreateAvatarUseCase usecaseInterface.CreateAvatarUseCaseInterface
	UpdateAvatarUseCase usecaseInterface.UpdateAvatarUseCaseInterface
	DeleteAvatarUseCase usecaseInterface.DeleteAvatarUseCaseInterface
}

func NewAvatarService() *AvatarService {
	return &AvatarService{
		GetAvatarUseCase:    usecase.NewGetAvatarUseCase(),
		GetAvatarsUseCase:   usecase.NewGetAvatarsUseCase(),
		CreateAvatarUseCase: usecase.NewCreateAvatarUseCase(),
		UpdateAvatarUseCase: usecase.NewUpdateAvatarUseCase(),
		DeleteAvatarUseCase: usecase.NewDeleteAvatarUseCase(),
	}
}

func (s *AvatarService) GetAvatar(input RequestGetAvatarInputDTO) (output usecase.GetAvatarOutputDTO, err error) {
	output, err = s.GetAvatarUseCase.Execute(usecase.GetAvatarInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar found")
	return
}

func (s *AvatarService) GetAvatars(input RequestGetAvatarsInputDTO) (output []usecase.GetAvatarsOutputDTO, err error) {
	output, err = s.GetAvatarsUseCase.Execute(usecase.GetAvatarsInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatars found")
	return
}

func (s *AvatarService) CreateAvatar(input RequestCreateAvatarInputDTO) (output usecase.CreateAvatarOutputDTO, err error) {
	output, err = s.CreateAvatarUseCase.Execute(usecase.CreateAvatarInputDTO{SVG: input.SVG})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(input RequestUpdateAvatarInputDTO) (output usecase.UpdateAvatarOutputDTO, err error) {
	output, err = s.UpdateAvatarUseCase.Execute(usecase.UpdateAvatarInputDTO{ID: input.ID, SVG: input.SVG})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(input RequestDeleteAvatarInputDTO) (output usecase.DeleteAvatarOutputDTO, err error) {
	output, err = s.DeleteAvatarUseCase.Execute(usecase.DeleteAvatarInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar deleted")
	return
}
