package response

type ResponseLogin struct {
	UserID                int32  `json:"user_id"`
	Token                 string `json:"token"`
	RefreshToken          string `json:"refresh_token"`
	Active                bool   `json:"active"`
	TokenExpiresIn        int32  `json:"token_expires_in"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
}

type ResponseRefreshToken struct {
	UserID                int32  `json:"user_id"`
	Token                 string `json:"token"`
	RefreshToken          string `json:"refresh_token"`
	Active                bool   `json:"active"`
	TokenExpiresIn        int32  `json:"token_expires_in"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
}
