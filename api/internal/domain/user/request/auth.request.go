package request

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	UserID       int32  `json:"user_id" validate:"number,required"`
	RefreshToken string `json:"refresh_token" validate:"jwt,required"`
}
