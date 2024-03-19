package request

type RequestCreateAvatarInputDTO struct {
	SVG string `json:"svg" validate:"required"`
}

type RequestUpdateAvatarInputDTO struct {
	ID  int32  `json:"id" validate:"number,required"`
	SVG string `json:"svg" validate:"required"`
}

type RequestGetAvatarInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetAvatarsInputDTO struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type RequestDeleteAvatarInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}
