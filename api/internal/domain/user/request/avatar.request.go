package request

type CreateAvatarRequest struct {
	SVG string `json:"svg" validate:"required"`
}

type UpdateAvatarRequest struct {
	ID  int32  `json:"id" validate:"number,required"`
	SVG string `json:"svg" validate:"required"`
}

type GetAvatarRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}

type GetAvatarsRequest struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type DeleteAvatarRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}
