package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/usecase"
)

type AvatarServiceInterface interface {
	CreateAvatar(body io.ReadCloser) (err error)
	GetAvatar(id int32) (output usecase.GetAvatarOutputDTO, err error)
	GetAvatars(limit int32, offset int32) (output []usecase.GetAvatarsOutputDTO, err error)
	UpdateAvatar(id int32, body io.ReadCloser) (err error)
	DeleteAvatar(id int32) (err error)
}

type AvatarService struct {
	repo repository.AvatarRepositoryInterface
}

func NewAvatarService(repo repository.AvatarRepositoryInterface) *AvatarService {
	return &AvatarService{
		repo: repo,
	}
}

func (s *AvatarService) GetAvatar(id int32) (output usecase.GetAvatarOutputDTO, err error) {

	input := usecase.GetAvatarInputDTO{
		ID: id,
	}

	output, err = usecase.NewGetAvatarUseCase(s.repo).Execute(input)
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

	output, err = usecase.NewGetAvatarsUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Avatars found")
	return
}

func (s *AvatarService) CreateAvatar(body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.CreateAvatarInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewCreateAvatarUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Avatar created")
	return
}

func (s *AvatarService) UpdateAvatar(id int32, body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdateAvatarInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	if err = usecase.NewUpdateAvatarUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Avatar updated")
	return
}

func (s *AvatarService) DeleteAvatar(id int32) (err error) {
	s.repo.Begin()

	input := usecase.DeleteAvatarInputDTO{
		ID: id,
	}

	if err = usecase.NewDeleteAvatarUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Avatar deleted")
	return
}
