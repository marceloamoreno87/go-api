package response

type GetUserResponse struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	AvatarID int32  `json:"avatar_id"`
	RoleID   int32  `json:"role_id"`
}

type GetUsersResponse struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	AvatarID int32  `json:"avatar_id"`
	RoleID   int32  `json:"role_id"`
}

type CreateUserResponse struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	AvatarID int32  `json:"avatar_id"`
	RoleID   int32  `json:"role_id"`
}

type UpdateUserResponse struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	AvatarID int32  `json:"avatar_id"`
	RoleID   int32  `json:"role_id"`
}

type DeleteUserResponse struct {
	ID int32 `json:"id"`
}

type UpdateUserPasswordResponse struct {
	ID int32 `json:"id"`
}

type VerifyUserResponse struct {
	ID int32 `json:"id"`
}

type ForgotPasswordResponse struct {
	ID int32 `json:"id"`
}
