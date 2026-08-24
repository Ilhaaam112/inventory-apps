package service

import (
	"errors"
	"strings"
	"time"

	"github.com/username/belajar_go/backend/internal/auth"
	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

// ErrKredensial sengaja satu untuk semua kegagalan login: username tidak
// ada, password salah, atau user tanpa role. Pesan yang berbeda-beda
// membocorkan informasi ke penyerang.
var (
	ErrKredensial = errors.New("username atau password salah")
	ErrSesi       = errors.New("sesi tidak valid, silakan login ulang")
)

type AuthService struct {
	users  *repository.UserRepository
	rbac   *repository.RBACRepository
	tokens *repository.RefreshTokenRepository
	mgr    *auth.Manager
}

func NewAuthService(
	users *repository.UserRepository,
	rbac *repository.RBACRepository,
	tokens *repository.RefreshTokenRepository,
	mgr *auth.Manager,
) *AuthService {
	return &AuthService{users: users, rbac: rbac, tokens: tokens, mgr: mgr}
}

// Hasil adalah bahan yang dipakai handler untuk menyusun response
// sekaligus memasang cookie.
type Hasil struct {
	Response     model.AuthResponse
	RefreshToken string
	RefreshExp   time.Time
}

func (s *AuthService) Login(username, password, userAgent, ip string) (Hasil, error) {
	username = strings.TrimSpace(username)
	if username == "" || password == "" {
		return Hasil{}, ErrKredensial
	}
	if len(username) > 100 || len(password) > 200 {
		return Hasil{}, ErrKredensial
	}

	user, err := s.users.GetByUsername(username)
	if err != nil {
		auth.BurnTime(password) // samakan waktu respons
		return Hasil{}, ErrKredensial
	}
	if !auth.CheckPassword(user.Password, password) {
		return Hasil{}, ErrKredensial
	}
	if user.RoleID == nil {
		return Hasil{}, ErrKredensial
	}

	return s.terbitkan(user, auth.NewID(), userAgent, ip)
}

// Refresh menjalankan rotasi: token lama ditandai terpakai, token baru
// diterbitkan dalam family yang sama. Kalau token yang sudah terpakai
// datang lagi, itu tanda token dicuri dan seluruh family dicabut.
func (s *AuthService) Refresh(rawToken, userAgent, ip string) (Hasil, error) {
	if rawToken == "" {
		return Hasil{}, ErrSesi
	}

	simpanan, err := s.tokens.FindByHash(auth.HashToken(rawToken))
	if err != nil {
		return Hasil{}, ErrSesi
	}

	if simpanan.Used {
		_ = s.tokens.RevokeFamily(simpanan.FamilyID)
		return Hasil{}, ErrSesi
	}
	if simpanan.Revoked || simpanan.Expired {
		return Hasil{}, ErrSesi
	}

	user, err := s.users.GetByID(simpanan.UserID)
	if err != nil || user.RoleID == nil {
		_ = s.tokens.RevokeFamily(simpanan.FamilyID)
		return Hasil{}, ErrSesi
	}

	if err := s.tokens.MarkUsed(simpanan.ID); err != nil {
		return Hasil{}, ErrSesi
	}
	return s.terbitkan(user, simpanan.FamilyID, userAgent, ip)
}

func (s *AuthService) Logout(rawToken string) {
	if rawToken == "" {
		return
	}
	if simpanan, err := s.tokens.FindByHash(auth.HashToken(rawToken)); err == nil {
		_ = s.tokens.RevokeFamily(simpanan.FamilyID)
	}
}

func (s *AuthService) terbitkan(user model.User, familyID, userAgent, ip string) (Hasil, error) {
	perms := s.rbac.Permissions(user.ID)

	accessToken, accessExp, err := s.mgr.IssueAccessToken(user, perms)
	if err != nil {
		return Hasil{}, err
	}

	rawRefresh, hashRefresh, err := auth.NewOpaqueToken()
	if err != nil {
		return Hasil{}, err
	}
	refreshExp := time.Now().Add(s.mgr.Config().RefreshTTL)
	if err := s.tokens.Create(user.ID, hashRefresh, familyID, refreshExp, userAgent, ip); err != nil {
		return Hasil{}, err
	}

	user.Password = ""
	user.Permissions = perms

	return Hasil{
		Response: model.AuthResponse{
			AccessToken: accessToken,
			TokenType:   "Bearer",
			ExpiresIn:   int(time.Until(accessExp).Seconds()),
			User:        user,
		},
		RefreshToken: rawRefresh,
		RefreshExp:   refreshExp,
	}, nil
}
