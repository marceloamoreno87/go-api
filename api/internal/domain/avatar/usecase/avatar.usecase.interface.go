package usecase

type CreateAvatarUseCaseInterface interface {
	Execute(input *CreateAvatarInputDTO) (err error)
}

type GetAvatarUseCaseInterface interface {
	Execute(id int32) (output *GetAvatarOutputDTO, err error)
}

type GetAvatarsUseCaseInterface interface {
	Execute(input GetAvatarsInputDTO) (output []*GetAvatarsOutputDTO, err error)
}

type UpdateAvatarUseCaseInterface interface {
	Execute(id int32, input *UpdateAvatarInputDTO) (err error)
}

type DeleteAvatarUseCaseInterface interface {
	Execute(id int32) (err error)
}
