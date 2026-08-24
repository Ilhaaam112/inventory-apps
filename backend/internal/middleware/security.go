package middleware

import (
	"log"
	"net/http"
)

// SecurityHeaders untuk API yang tidak pernah merender HTML.
func SecurityHeaders(production bool) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := w.Header()
			h.Set("X-Content-Type-Options", "nosniff")
			h.Set("X-Frame-Options", "DENY")
			h.Set("Referrer-Policy", "no-referrer")
			h.Set("Cross-Origin-Resource-Policy", "same-site")
			h.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
			h.Set("Cache-Control", "no-store")
			h.Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

			// HSTS hanya bermakna di HTTPS, dan berbahaya kalau
			// dipasang saat development di http://localhost.
			if production {
				h.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Recover menahan panic supaya server tidak mati, dan supaya stack
// trace berhenti di log server, tidak sampai ke browser.
func Recover(production bool) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("PANIC %s %s: %v", r.Method, r.URL.Path, rec)
					WriteError(w, http.StatusInternalServerError, "terjadi kesalahan pada server")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
