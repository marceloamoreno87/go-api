package usecase

type GetJWTUseCaseInterface interface {
	Execute(input GetJWTInputDTO) (output GetJWTOutputDTO, err error)
}

type GetRefreshJWTUseCaseInterface interface {
	Execute(input GetRefreshJWTInputDTO) (output GetRefreshJWTOutputDTO, err error)
}
