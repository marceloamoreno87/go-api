package response

type ResponseGetAvatar struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type ResponseGetAvatars struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type ResponseCreateAvatar struct {
	SVG string `json:"svg"`
}

type ResponseUpdateAvatar struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type ResponseDeleteAvatar struct {
	ID int32 `json:"id"`
}
