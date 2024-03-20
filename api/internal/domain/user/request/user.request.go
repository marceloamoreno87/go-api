package request

type RequestVerifyUser struct {
	Hash string `json:"hash" validate:"required"`
}

type RequestUpdateUserPassword struct {
	Hash     string `json:"hash" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestForgotPassword struct {
	Email string `json:"email" validate:"required"`
}

type RequestCreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type RequestUpdateUser struct {
	ID    int32  `json:"id" validate:"number,required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type RequestDeleteUser struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetUser struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetUsers struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}
