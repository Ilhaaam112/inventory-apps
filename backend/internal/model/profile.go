package model

type UpdateProfileRequest struct {
	ID          int    `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
}

type ChangePasswordRequest struct {
	ID          int    `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}