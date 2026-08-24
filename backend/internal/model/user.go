package model

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"-"`
	NamaLengkap  string `json:"nama_lengkap"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}