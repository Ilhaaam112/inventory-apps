package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

// CORS hanya memantulkan origin yang persis ada di daftar putih.
// Tidak pernah memakai "*", karena API ini mengirim credentials
// (cookie refresh token) dan browser memang melarang kombinasi itu.
func CORS(allowedOrigins []string) Middleware {
	izin := map[string]bool{}
	for _, o := range allowedOrigins {
		izin[strings.ToLower(strings.TrimRight(o, "/"))] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Vary wajib, supaya cache tidak menyajikan header CORS
			// milik origin lain.
			w.Header().Add("Vary", "Origin")

			if origin != "" && izin[strings.ToLower(strings.TrimRight(origin, "/"))] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.Header().Set("Access-Control-Max-Age", strconv.Itoa(600))
			} else if origin != "" {
				// Origin asing: preflight langsung ditolak.
				if r.Method == http.MethodOptions {
					w.WriteHeader(http.StatusForbidden)
					return
				}
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
