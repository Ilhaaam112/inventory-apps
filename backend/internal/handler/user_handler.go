package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
		return
	}

	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"data tidak valid"}`, http.StatusBadRequest)
		return
	}

	user, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusUnauthorized)
		return
	}

	response := model.LoginResponse{
		Message: "Login berhasil",
		User:    user,
	}
	json.NewEncoder(w).Encode(response)
}






func (h *UserHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, `{"error":"id tidak valid"}`, http.StatusBadRequest)
			return
		}
		user, err := h.service.GetProfile(id)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)

	case http.MethodPut:
		var req model.UpdateProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"body tidak valid"}`, http.StatusBadRequest)
			return
		}
		user, err := h.service.UpdateProfile(req.ID, req.NamaLengkap)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
		return
	}
	var req model.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"body tidak valid"}`, http.StatusBadRequest)
		return
	}
	if err := h.service.ChangePassword(req.ID, req.OldPassword, req.NewPassword); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"message":"password berhasil diubah"}`))
}