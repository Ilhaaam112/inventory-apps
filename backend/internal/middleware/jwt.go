package middleware

import (
	"net/http"
	"strings"

	"github.com/username/belajar_go/backend/internal/auth"
)

const pesanTidakTerautentikasi = "tidak terautentikasi"

// RequireAuth memverifikasi access token lalu menaruh identitas user
// di context. Gagal apa pun alasannya -> 401 dengan pesan yang sama.
func RequireAuth(mgr *auth.Manager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				WriteError(w, http.StatusUnauthorized, pesanTidakTerautentikasi)
				return
			}

			bagian := strings.SplitN(header, " ", 2)
			if len(bagian) != 2 || !strings.EqualFold(bagian[0], "Bearer") {
				WriteError(w, http.StatusUnauthorized, pesanTidakTerautentikasi)
				return
			}

			claims, err := mgr.ParseAccessToken(strings.TrimSpace(bagian[1]))
			if err != nil {
				WriteError(w, http.StatusUnauthorized, pesanTidakTerautentikasi)
				return
			}

			next.ServeHTTP(w, withUser(r, AuthUser{
				ID:          claims.UserID,
				Username:    claims.Username,
				Role:        claims.Role,
				Permissions: claims.Permissions,
			}))
		})
	}
}
