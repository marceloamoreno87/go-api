package request

type RequestVerifyUserInputDTO struct {
	Hash string `json:"hash" validate:"required"`
}

type RequestUpdateUserPasswordInputDTO struct {
	Hash     string `json:"hash" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestForgotPasswordInputDTO struct {
	Email string `json:"email" validate:"required"`
}

type RequestCreateUserInputDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type RequestUpdateUserInputDTO struct {
	ID    int32  `json:"id" validate:"number,required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type RequestDeleteUserInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetUserInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetUsersInputDTO struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}
