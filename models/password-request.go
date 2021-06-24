package models

type PasswordRequest struct {
	ID          string `json:"id"`
	NewPassword string `json:"newPassword"`
}
