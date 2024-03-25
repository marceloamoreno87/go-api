package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
	"github.com/marceloamoreno/goapi/internal/shared/response"
)

type AvatarHandler struct {
	response.Responses
	service service.AvatarService
}

func NewAvatarHandler() *AvatarHandler {
	return &AvatarHandler{
		service: *service.NewAvatarService(),
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
// @Router /avatar/{id} [get]
// @Security     JWT
func (h *AvatarHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	input := request.RequestGetAvatar{
		ID: helper.GetID(r),
	}
	output, err := h.service.GetAvatar(r.Context(), input)
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
// @Router /avatar [get]
// @Security     JWT
func (h *AvatarHandler) GetAvatars(w http.ResponseWriter, r *http.Request) {
	limit, offset := helper.GetLimitAndOffset(r)
	input := request.RequestGetAvatars{
		Limit:  limit,
		Offset: offset,
	}
	output, err := h.service.GetAvatars(r.Context(), input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// CreateAvatar godoc
// @Summary Create Avatar
// @Description Create Avatar
// @Tags Avatar
// @Accept  json
// @Produce  json
// @Param avatar body request.RequestCreateAvatar true "Avatar"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /avatar [post]
// @Security     JWT
func (h *AvatarHandler) CreateAvatar(w http.ResponseWriter, r *http.Request) {
	input := request.RequestCreateAvatar{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err := h.service.CreateAvatar(r.Context(), input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}

// UpdateAvatar godoc
// @Summary Update Avatar
// @Description Update Avatar
// @Tags Avatar
// @Accept  json
// @Produce  json
// @Param id path string true "Avatar ID"
// @Param avatar body request.RequestUpdateAvatar true "Avatar"
// @Success 200 {object} response.Response{data=nil}
// @Failure 400 {object} response.ResponseError{}
// @Router /avatar/{id} [put]
// @Security     JWT
func (h *AvatarHandler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	input := request.RequestUpdateAvatar{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err := h.service.UpdateAvatar(r.Context(), input)
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
// @Router /avatar/{id} [delete]
// @Security     JWT
func (h *AvatarHandler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	input := request.RequestDeleteAvatar{
		ID: helper.GetID(r),
	}
	output, err := h.service.DeleteAvatar(r.Context(), input)
	if err != nil {
		slog.Info("err", err)
		h.SendResponseError(w, h.NewResponseError(err.Error()))
		return
	}
	h.SendResponse(w, h.NewResponse(output))
}
