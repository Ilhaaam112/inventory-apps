package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/username/belajar_go/backend/internal/auth"
	"github.com/username/belajar_go/backend/internal/middleware"
	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

// Cookie refresh token dibatasi Path ke endpoint auth saja, jadi dia
// tidak ikut terkirim pada puluhan request biasa lainnya.
const (
	namaCookie = "refresh_token"
	pathCookie = "/api/v1/auth"
)

type AuthHandler struct {
	service *service.AuthService
	cfg     auth.Config
}

func NewAuthHandler(s *service.AuthService, cfg auth.Config) *AuthHandler {
	return &AuthHandler{service: s, cfg: cfg}
}

// POST /api/v1/auth/login
func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}

	var in model.LoginInput
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "data tidak valid")
		return
	}

	hasil, err := h.service.Login(in.Username, in.Password, r.UserAgent(), h.ip(r))
	if err != nil {
		writeError(w, http.StatusUnauthorized, pesanAman(err))
		return
	}

	h.pasangCookie(w, hasil.RefreshToken, hasil.RefreshExp)
	writeJSON(w, hasil.Response)
}

// POST /api/v1/auth/refresh
func (h *AuthHandler) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}

	cookie, err := r.Cookie(namaCookie)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "sesi tidak valid, silakan login ulang")
		return
	}

	hasil, err := h.service.Refresh(cookie.Value, r.UserAgent(), h.ip(r))
	if err != nil {
		h.hapusCookie(w)
		writeError(w, http.StatusUnauthorized, pesanAman(err))
		return
	}

	h.pasangCookie(w, hasil.RefreshToken, hasil.RefreshExp)
	writeJSON(w, hasil.Response)
}

// POST /api/v1/auth/logout
func (h *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}

	if cookie, err := r.Cookie(namaCookie); err == nil {
		h.service.Logout(cookie.Value)
	}
	h.hapusCookie(w)
	writeJSON(w, map[string]string{"message": "berhasil keluar"})
}

func (h *AuthHandler) ip(r *http.Request) string {
	return middleware.ClientIP(r, h.cfg.Production)
}

func (h *AuthHandler) pasangCookie(w http.ResponseWriter, nilai string, exp time.Time) {
	sameSite := http.SameSiteLaxMode
	if h.cfg.Production {
		sameSite = http.SameSiteStrictMode
	}
	http.SetCookie(w, &http.Cookie{
		Name:     namaCookie,
		Value:    nilai,
		Path:     pathCookie,
		Domain:   h.cfg.CookieDomain,
		Expires:  exp,
		MaxAge:   int(time.Until(exp).Seconds()),
		HttpOnly: true,             // tidak bisa dibaca JavaScript
		Secure:   h.cfg.Production, // di dev http://localhost harus false
		SameSite: sameSite,
	})
}

func (h *AuthHandler) hapusCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     namaCookie,
		Value:    "",
		Path:     pathCookie,
		Domain:   h.cfg.CookieDomain,
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   h.cfg.Production,
		SameSite: http.SameSiteLaxMode,
	})
}

// pesanAman hanya meneruskan error yang memang dirancang dibaca user.
// Error lain (database, bcrypt) diganti pesan generik.
func pesanAman(err error) string {
	if errors.Is(err, service.ErrKredensial) || errors.Is(err, service.ErrSesi) {
		return err.Error()
	}
	return "terjadi kesalahan pada server"
}
