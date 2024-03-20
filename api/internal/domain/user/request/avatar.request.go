package request

type RequestCreateAvatar struct {
	SVG string `json:"svg" validate:"required"`
}

type RequestUpdateAvatar struct {
	ID  int32  `json:"id" validate:"number,required"`
	SVG string `json:"svg" validate:"required"`
}

type RequestGetAvatar struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetAvatars struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type RequestDeleteAvatar struct {
	ID int32 `json:"id" validate:"number,required"`
}
