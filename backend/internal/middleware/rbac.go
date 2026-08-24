package middleware

import "net/http"

const pesanTidakBerhak = "tidak memiliki izin untuk aksi ini"

// RequirePermission dipakai untuk route yang izinnya sama untuk
// semua HTTP method.
func RequirePermission(code string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, ok := UserFrom(r)
			if !ok {
				WriteError(w, http.StatusUnauthorized, pesanTidakTerautentikasi)
				return
			}
			if !u.Can(code) {
				WriteError(w, http.StatusForbidden, pesanTidakBerhak)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// RequirePerMethod memetakan HTTP method ke permission, karena satu
// handler di project ini melayani beberapa method sekaligus.
// Method yang tidak ada di peta akan ditolak 403.
func RequirePerMethod(peta map[string]string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, ok := UserFrom(r)
			if !ok {
				WriteError(w, http.StatusUnauthorized, pesanTidakTerautentikasi)
				return
			}
			code, ada := peta[r.Method]
			if !ada || !u.Can(code) {
				WriteError(w, http.StatusForbidden, pesanTidakBerhak)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
