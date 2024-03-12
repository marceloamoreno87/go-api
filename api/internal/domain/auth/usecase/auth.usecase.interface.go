package usecase

type LoginUseCaseInterface interface {
	Execute(input LoginInputDTO) (err error)
}
