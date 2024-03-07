package handler

import (
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/avatar/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/avatar/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type AvatarHandlerInterface interface {
	GetAvatar(w http.ResponseWriter, r *http.Request)
	GetAvatars(w http.ResponseWriter, r *http.Request)
	CreateAvatar(w http.ResponseWriter, r *http.Request)
	UpdateAvatar(w http.ResponseWriter, r *http.Request)
	DeleteAvatar(w http.ResponseWriter, r *http.Request)
}

type AvatarHandler struct {
	response.Responses
	service service.AvatarServiceInterface
}

func NewAvatarHandler(
	service service.AvatarServiceInterface,
) *AvatarHandler {
	return &AvatarHandler{
		service: service,
	}
}

// GetAvatar godoc
// @Summary Get Avatar
// @Description Get Avatar
// @Tags Avatar
// @Accept  json
// @Produce  json
// @Param id path string true "Avatar ID"
// @Success 200 {object} response.Response{data=usecase.GetAvatarOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /Avatar/{id} [get]
// @Security     JWT
func (h *AvatarHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.GetAvatar(helper.GetID(r))
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// GetAvatars godoc
// @Summary Get Avatars
// @Description Get Avatars
// @Tags Avatar
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} response.Response{data=[]usecase.GetAvatarsOutputDTO}
// @Failure 400 {object} response.ResponseError{}
// @Router /Avatar [get]
// @Security     JWT
func (h *AvatarHandler) GetAvatars(w http.ResponseWriter, r *http.Request) {
	limit, offset := helper.GetLimitAndOffset(r)
	output, err := h.service.GetAvatars(limit, offset)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// CreateRole godoc
// @Summary Create Avatar
// @Description Create Avatar
// @Tags Avatar
// @Accept  json
// @Produce  json
// @Param role body usecase.CreateAvatarInputDTO true "Avatar"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /role [post]
// @Security     JWT
func (h *AvatarHandler) CreateAvatar(w http.ResponseWriter, r *http.Request) {
	if err := h.service.CreateAvatar(r.Body); err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}

// UpdateRole godoc
// @Summary Update Avatar
// @Description Update Avatar
// @Tags Avatar
// @Accept  json
// @Produce  json
// @Param id path string true "Avatar ID"
// @Param role body usecase.UpdateAvatarInputDTO true "Avatar"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /Avatar/{id} [put]
// @Security     JWT
func (h *AvatarHandler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	if err := h.service.UpdateAvatar(helper.GetID(r), r.Body); err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}

// DeleteAvatar godoc
// @Summary Delete Avatar
// @Description Delete Avatar
// @Tags Avatar
// @Accept  json
// @Produce  json
// @Param id path string true "Avatar ID"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Security ApiKeyAuth
// @Router /Avatar/{id} [delete]
// @Security     JWT
func (h *AvatarHandler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeleteAvatar(helper.GetID(r)); err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(nil))
}
