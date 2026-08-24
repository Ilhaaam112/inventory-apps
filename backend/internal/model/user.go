package model

type User struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Password    string   `json:"-"`
	NamaLengkap string   `json:"nama_lengkap"`
	RoleID      *int     `json:"role_id,omitempty"`
	RoleName    string   `json:"role"`
	Permissions []string `json:"permissions,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type UpdateProfileRequest struct {
	ID          int    `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
}

type ChangePasswordRequest struct {
	ID          int    `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
