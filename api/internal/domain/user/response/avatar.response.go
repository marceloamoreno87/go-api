package response

type GetAvatarResponse struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type GetAvatarsResponse struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type CreateAvatarResponse struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type UpdateAvatarResponse struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type DeleteAvatarResponse struct {
	ID int32 `json:"id"`
}
