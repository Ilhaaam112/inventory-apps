package model

// LoginInput divalidasi ulang di backend, tidak bergantung pada React.
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse adalah balasan login dan refresh.
// Refresh token TIDAK ada di sini: dia dikirim lewat HttpOnly cookie
// supaya tidak bisa dibaca JavaScript.
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	User        User   `json:"user"`
}

// RefreshToken sengaja memakai boolean, bukan time.Time. Status
// kadaluarsa dihitung di sisi MySQL, jadi kode ini tidak bergantung
// pada opsi parseTime di DSN database.
type RefreshToken struct {
	ID       int64
	UserID   int
	FamilyID string
	Expired  bool
	Used     bool
	Revoked  bool
}
