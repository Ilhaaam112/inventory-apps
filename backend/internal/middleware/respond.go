package middleware

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

// WriteError selalu memakai pesan yang sudah disiapkan, tidak pernah
// meneruskan error database atau stack trace ke client.
func WriteError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// ClientIP mengambil IP asli. X-Forwarded-For hanya dipercaya kalau
// server memang berada di belakang reverse proxy yang kamu kendalikan.
func ClientIP(r *http.Request, trustProxy bool) string {
	if trustProxy {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			if i := strings.Index(xff, ","); i > 0 {
				return strings.TrimSpace(xff[:i])
			}
			return strings.TrimSpace(xff)
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
