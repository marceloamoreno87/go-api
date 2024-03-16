package service

import (
	"encoding/json"
	"io"
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AvatarService struct {
	NewGetAvatarUseCase    usecaseInterface.GetAvatarUseCaseInterface
	NewGetAvatarsUseCase   usecaseInterface.GetAvatarsUseCaseInterface
	NewCreateAvatarUseCase usecaseInterface.CreateAvatarUseCaseInterface
	NewUpdateAvatarUseCase usecaseInterface.UpdateAvatarUseCaseInterface
	NewDeleteAvatarUseCase usecaseInterface.DeleteAvatarUseCaseInterface
}

func NewAvatarService() *AvatarService {
	return &AvatarService{
		NewGetAvatarUseCase:    usecase.NewGetAvatarUseCase(),
		NewGetAvatarsUseCase:   usecase.NewGetAvatarsUseCase(),
		NewCreateAvatarUseCase: usecase.NewCreateAvatarUseCase(),
		NewUpdateAvatarUseCase: usecase.NewUpdateAvatarUseCase(),
		NewDeleteAvatarUseCase: usecase.NewDeleteAvatarUseCase(),
	}
}

func (s *AvatarService) GetAvatar(id int32) (output usecase.GetAvatarOutputDTO, err error) {
	input := usecase.GetAvatarInputDTO{
		ID: id,
	}

	output, err = s.NewGetAvatarUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar found")
	return
}

func (s *AvatarService) GetAvatars(limit int32, offset int32) (output []usecase.GetAvatarsOutputDTO, err error) {
	input := usecase.GetAvatarsInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	output, err = s.NewGetAvatarsUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatars found")
	return
}

func (s *AvatarService) CreateAvatar(body io.ReadCloser) (output usecase.CreateAvatarOutputDTO, err error) {
	input := usecase.CreateAvatarInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.NewCreateAvatarUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(id int32, body io.ReadCloser) (output usecase.UpdateAvatarOutputDTO, err error) {
	input := usecase.UpdateAvatarInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = s.NewUpdateAvatarUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(id int32) (output usecase.DeleteAvatarOutputDTO, err error) {
	input := usecase.DeleteAvatarInputDTO{
		ID: id,
	}

	output, err = s.NewDeleteAvatarUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatar deleted")
	return
}
