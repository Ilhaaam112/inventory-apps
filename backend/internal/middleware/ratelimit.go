package middleware

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type pengunjung struct {
	limiter  *rate.Limiter
	terakhir time.Time
}

// RateLimit membatasi permintaan per alamat IP memakai token bucket.
//
// Catatan penting: penyimpanannya di memori proses ini saja. Kalau
// server direstart, hitungannya hilang; kalau nanti dijalankan lebih
// dari satu instance, tiap instance punya hitungan sendiri. Untuk
// produksi ber-scale, pindahkan ke Redis.
func RateLimit(setiap time.Duration, burst int, trustProxy bool) Middleware {
	var mu sync.Mutex
	daftar := map[string]*pengunjung{}

	// Bersihkan entri lama supaya map tidak tumbuh tanpa batas.
	go func() {
		for range time.Tick(3 * time.Minute) {
			mu.Lock()
			for ip, v := range daftar {
				if time.Since(v.terakhir) > 10*time.Minute {
					delete(daftar, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := ClientIP(r, trustProxy)

			mu.Lock()
			v, ada := daftar[ip]
			if !ada {
				v = &pengunjung{limiter: rate.NewLimiter(rate.Every(setiap), burst)}
				daftar[ip] = v
			}
			v.terakhir = time.Now()
			izin := v.limiter.Allow()
			mu.Unlock()

			if !izin {
				w.Header().Set("Retry-After", "60")
				WriteError(w, http.StatusTooManyRequests, "terlalu banyak permintaan, coba lagi sebentar")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
