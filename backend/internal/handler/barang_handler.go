package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type BarangHandler struct {
	service *service.BarangService
}

func NewBarangHandler(s *service.BarangService) *BarangHandler {
	return &BarangHandler{service: s}
}

func (h *BarangHandler) HandleBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		data := h.service.GetAll()
		json.NewEncoder(w).Encode(data)

	case http.MethodPost:
		var b model.Barang
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, `{"error":"data tidak valid"}`, http.StatusBadRequest)
			return
		}
		created, err := h.service.Create(b)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}

func (h *BarangHandler) HandleBarangByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/barang/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"id tidak valid"}`, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		b, err := h.service.GetByID(id)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(b)

	case http.MethodPut:
		var b model.Barang
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, `{"error":"data tidak valid"}`, http.StatusBadRequest)
			return
		}
		updated, err := h.service.Update(id, b)
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
		json.NewEncoder(w).Encode(map[string]string{"message": "barang berhasil dihapus"})

	default:
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}