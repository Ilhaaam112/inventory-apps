package middleware

import "net/http"

// Middleware membungkus satu handler menjadi handler baru.
type Middleware func(http.Handler) http.Handler

// Chain menyusun middleware supaya urutan penulisan sama dengan
// urutan eksekusi:
//
//	Chain(A, B, C)(h)  ->  A( B( C(h) ) )
func Chain(mws ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(mws) - 1; i >= 0; i-- {
			final = mws[i](final)
		}
		return final
	}
}
