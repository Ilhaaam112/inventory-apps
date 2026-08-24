package auth

import (
	"errors"
	"os"
	"strings"
	"time"
)

// Config dibaca sekali saat server start. Kalau JWT_SECRET tidak ada
// atau terlalu pendek, server sengaja menolak jalan: lebih baik gagal
// keras daripada jalan dengan kunci lemah.
type Config struct {
	Secret         []byte
	Issuer         string
	Audience       string
	AccessTTL      time.Duration
	RefreshTTL     time.Duration
	Production     bool
	CookieDomain   string
	AllowedOrigins []string
}

func LoadConfig() (Config, error) {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 32 {
		return Config{}, errors.New("JWT_SECRET wajib diisi dan minimal 32 karakter (lihat backend/.env.example)")
	}

	cfg := Config{
		Secret:       []byte(secret),
		Issuer:       envOr("JWT_ISSUER", "belajar-go-inventory"),
		Audience:     envOr("JWT_AUDIENCE", "belajar-go-web"),
		AccessTTL:    durationOr("ACCESS_TOKEN_TTL", 15*time.Minute),
		RefreshTTL:   durationOr("REFRESH_TOKEN_TTL", 7*24*time.Hour),
		Production:   strings.EqualFold(os.Getenv("APP_ENV"), "production"),
		CookieDomain: os.Getenv("COOKIE_DOMAIN"),
	}

	for _, o := range strings.Split(envOr("CORS_ALLOWED_ORIGINS", "http://localhost:5173"), ",") {
		if o = strings.TrimSpace(o); o != "" {
			cfg.AllowedOrigins = append(cfg.AllowedOrigins, o)
		}
	}
	if len(cfg.AllowedOrigins) == 0 {
		return Config{}, errors.New("CORS_ALLOWED_ORIGINS tidak boleh kosong")
	}
	return cfg, nil
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func durationOr(key string, fallback time.Duration) time.Duration {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		if d, err := time.ParseDuration(v); err == nil && d > 0 {
			return d
		}
	}
	return fallback
}
