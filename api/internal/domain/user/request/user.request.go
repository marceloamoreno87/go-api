package request

type VerifyUserRequest struct {
	Hash string `json:"hash" validate:"required"`
}

type UpdateUserPasswordRequest struct {
	Hash     string `json:"hash" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	ID    int32  `json:"id" validate:"number,required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type DeleteUserRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}

type GetUserRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}

type GetUsersRequest struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}
