package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type KategoriHandler struct {
	service *service.KategoriService
}

func NewKategoriHandler(s *service.KategoriService) *KategoriHandler {
	return &KategoriHandler{service: s}
}

func (h *KategoriHandler) HandleKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.service.GetAll())

	case http.MethodPost:
		var k model.Kategori
		if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
			http.Error(w, `{"error":"data tidak valid"}`, http.StatusBadRequest)
			return
		}
		created, err := h.service.Create(k)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}

func (h *KategoriHandler) HandleKategoriByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"id tidak valid"}`, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		k, err := h.service.GetByID(id)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(k)

	case http.MethodPut:
		var k model.Kategori
		if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
			http.Error(w, `{"error":"data tidak valid"}`, http.StatusBadRequest)
			return
		}
		updated, err := h.service.Update(id, k)
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
		json.NewEncoder(w).Encode(map[string]string{"message": "kategori berhasil dihapus"})

	default:
		http.Error(w, `{"error":"method tidak diizinkan"}`, http.StatusMethodNotAllowed)
	}
}