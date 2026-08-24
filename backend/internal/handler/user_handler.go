package handler

import (
	"net/http"

	"github.com/username/belajar_go/backend/internal/middleware"
	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// HandleProfile sekarang selalu memakai id dari access token.
// Parameter ?id= diabaikan, supaya user tidak bisa membuka atau
// mengubah profil orang lain hanya dengan mengganti angka di URL.
func (h *UserHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	id := middleware.UserID(r)
	if id == 0 {
		writeError(w, http.StatusUnauthorized, "tidak terautentikasi")
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := h.service.GetProfile(id)
		if err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, user)

	case http.MethodPut:
		var req model.UpdateProfileRequest
		if err := decodeJSON(r, &req); err != nil {
			writeError(w, http.StatusBadRequest, "data tidak valid")
			return
		}
		user, err := h.service.UpdateProfile(id, req.NamaLengkap)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, user)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}

func (h *UserHandler) HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	id := middleware.UserID(r)
	if id == 0 {
		writeError(w, http.StatusUnauthorized, "tidak terautentikasi")
		return
	}

	var req model.ChangePasswordRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "data tidak valid")
		return
	}
	if err := h.service.ChangePassword(id, req.OldPassword, req.NewPassword); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, map[string]string{"message": "password berhasil diubah"})
}
