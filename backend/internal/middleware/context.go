package middleware

import (
	"context"
	"net/http"
)

type ctxKey int

const userKey ctxKey = iota

// AuthUser adalah identitas yang sudah terverifikasi dari access token.
// Handler membaca ini, BUKAN body request. Inilah yang menutup celah
// user_id yang dulu dikirim dari React.
type AuthUser struct {
	ID          int
	Username    string
	Role        string
	Permissions []string
}

func (u AuthUser) Can(code string) bool {
	for _, p := range u.Permissions {
		if p == code {
			return true
		}
	}
	return false
}

func withUser(r *http.Request, u AuthUser) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userKey, u))
}

// UserFrom dipakai handler untuk mengambil user yang sedang login.
func UserFrom(r *http.Request) (AuthUser, bool) {
	u, ok := r.Context().Value(userKey).(AuthUser)
	return u, ok
}

// UserID mengembalikan 0 kalau route-nya publik.
func UserID(r *http.Request) int {
	if u, ok := UserFrom(r); ok {
		return u.ID
	}
	return 0
}
