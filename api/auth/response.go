package auth

type TokenResponse struct {
	Expired      int64  `json:"expired"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
