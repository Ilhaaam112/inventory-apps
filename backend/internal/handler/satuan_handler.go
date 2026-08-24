package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type SatuanHandler struct {
	service *service.SatuanService
}

func NewSatuanHandler(s *service.SatuanService) *SatuanHandler {
	return &SatuanHandler{service: s}
}

func (h *SatuanHandler) HandleSatuan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.service.GetAll())

	case http.MethodPost:
		var m model.Satuan
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, `{"error":"data tidak valid"}`, http.StatusBadRequest)
			return
		}
		created, err := h.service.Create(m)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}

func (h *SatuanHandler) HandleSatuanByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/satuan/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"id tidak valid"}`, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		m, err := h.service.GetByID(id)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(m)

	case http.MethodPut:
		var m model.Satuan
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, `{"error":"data tidak valid"}`, http.StatusBadRequest)
			return
		}
		updated, err := h.service.Update(id, m)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "satuan berhasil dihapus"})

	default:
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}
