package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/username/belajar_go/backend/internal/model"
)

var ErrRefreshNotFound = errors.New("refresh token tidak ditemukan")

type RefreshTokenRepository struct {
	db *sql.DB
}

func NewRefreshTokenRepository(db *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(userID int, tokenHash, familyID string, expiresAt time.Time, ua, ip string) error {
	if len(ua) > 255 {
		ua = ua[:255]
	}
	if len(ip) > 45 {
		ip = ip[:45]
	}
	_, err := r.db.Exec(`
		INSERT INTO refresh_tokens (user_id, token_hash, family_id, user_agent, ip, expires_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		userID, tokenHash, familyID, ua, ip, expiresAt)
	return err
}

// FindByHash menghitung status kadaluarsa langsung di MySQL, supaya
// tidak ada kolom DATETIME yang perlu di-scan ke time.Time.
func (r *RefreshTokenRepository) FindByHash(tokenHash string) (model.RefreshToken, error) {
	var t model.RefreshToken
	err := r.db.QueryRow(`
		SELECT id, user_id, family_id,
		       (expires_at < NOW())       AS kadaluarsa,
		       (used_at    IS NOT NULL)   AS terpakai,
		       (revoked_at IS NOT NULL)   AS dicabut
		FROM refresh_tokens
		WHERE token_hash = ?`, tokenHash,
	).Scan(&t.ID, &t.UserID, &t.FamilyID, &t.Expired, &t.Used, &t.Revoked)
	if err != nil {
		return model.RefreshToken{}, ErrRefreshNotFound
	}
	return t, nil
}

func (r *RefreshTokenRepository) MarkUsed(id int64) error {
	res, err := r.db.Exec(
		"UPDATE refresh_tokens SET used_at = NOW() WHERE id = ? AND used_at IS NULL", id)
	if err != nil {
		return err
	}
	// Kalau tidak ada baris yang berubah, berarti ada request lain yang
	// memakai token yang sama lebih dulu. Perlakukan sebagai gagal.
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("token sudah dipakai")
	}
	return nil
}

// RevokeFamily mencabut seluruh token satu rantai rotasi sekaligus.
// Dipanggil saat logout, dan saat terdeteksi token dipakai ulang.
func (r *RefreshTokenRepository) RevokeFamily(familyID string) error {
	_, err := r.db.Exec(
		"UPDATE refresh_tokens SET revoked_at = NOW() WHERE family_id = ? AND revoked_at IS NULL",
		familyID)
	return err
}

func (r *RefreshTokenRepository) RevokeAllForUser(userID int) error {
	_, err := r.db.Exec(
		"UPDATE refresh_tokens SET revoked_at = NOW() WHERE user_id = ? AND revoked_at IS NULL",
		userID)
	return err
}

// DeleteExpired dipanggil berkala dari main.go.
func (r *RefreshTokenRepository) DeleteExpired() error {
	_, err := r.db.Exec("DELETE FROM refresh_tokens WHERE expires_at < NOW() - INTERVAL 30 DAY")
	return err
}
