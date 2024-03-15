package handler

import (
	"log/slog"
	"net/http"

	serviceInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	_ "github.com/marceloamoreno/goapi/internal/domain/user/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type AvatarHandler struct {
	response.Responses
	service serviceInterface.AvatarServiceInterface
}

func NewAvatarHandler() *AvatarHandler {
	return &AvatarHandler{
		service: service.NewAvatarService(),
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
	output, err := h.service.CreateAvatar(r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
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
	output, err := h.service.UpdateAvatar(helper.GetID(r), r.Body)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
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
	output, err := h.service.DeleteAvatar(helper.GetID(r))
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
