package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type RenewTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
type ChangePassRequest struct {
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
