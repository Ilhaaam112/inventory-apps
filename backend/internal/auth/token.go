package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/username/belajar_go/backend/internal/model"
)

const (
	TokenTypeAccess = "access"
)

var ErrTokenInvalid = errors.New("token tidak valid")

// Claims adalah isi access token. Field pendek supaya token tidak gemuk.
type Claims struct {
	UserID      int      `json:"uid"`
	Username    string   `json:"usr"`
	Role        string   `json:"role"`
	Permissions []string `json:"perms"`
	TokenType   string   `json:"typ"`
	jwt.RegisteredClaims
}

type Manager struct {
	cfg Config
}

func NewManager(cfg Config) *Manager {
	return &Manager{cfg: cfg}
}

func (m *Manager) Config() Config { return m.cfg }

func (m *Manager) IssueAccessToken(u model.User, perms []string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(m.cfg.AccessTTL)

	claims := Claims{
		UserID:      u.ID,
		Username:    u.Username,
		Role:        u.RoleName,
		Permissions: perms,
		TokenType:   TokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.cfg.Issuer,
			Audience:  jwt.ClaimStrings{m.cfg.Audience},
			Subject:   strconv.Itoa(u.ID),
			ID:        NewID(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.cfg.Secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, exp, nil
}

// ParseAccessToken memeriksa, berurutan:
//   1. algoritma tanda tangan  -> hanya HS256 (menolak serangan alg=none)
//   2. tanda tangan            -> lewat secret
//   3. exp dan nbf             -> lewat validator bawaan
//   4. issuer                  -> harus sama dengan konfigurasi
//   5. audience                -> harus sama dengan konfigurasi
//   6. tipe token              -> harus "access", bukan refresh
//
// Semua kegagalan dikembalikan sebagai satu error yang sama, supaya
// penyerang tidak bisa membedakan "expired" dari "tanda tangan salah".
func (m *Manager) ParseAccessToken(raw string) (*Claims, error) {
	var c Claims

	_, err := jwt.ParseWithClaims(raw, &c,
		func(t *jwt.Token) (interface{}, error) { return m.cfg.Secret, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(m.cfg.Issuer),
		jwt.WithAudience(m.cfg.Audience),
	)
	if err != nil {
		return nil, ErrTokenInvalid
	}
	if c.ExpiresAt == nil {
		return nil, ErrTokenInvalid
	}
	if c.TokenType != TokenTypeAccess {
		return nil, ErrTokenInvalid
	}
	if c.UserID == 0 {
		return nil, ErrTokenInvalid
	}
	return &c, nil
}
