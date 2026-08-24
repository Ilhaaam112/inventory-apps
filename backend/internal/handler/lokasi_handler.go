package handler

import (
	"net/http"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type LokasiHandler struct {
	service *service.LokasiService
}

func NewLokasiHandler(s *service.LokasiService) *LokasiHandler {
	return &LokasiHandler{service: s}
}

func (h *LokasiHandler) HandleLokasi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, h.service.GetAll())

	case http.MethodPost:
		var m model.Lokasi
		if err := decodeJSON(r, &m); err != nil {
			writeError(w, http.StatusBadRequest, "data tidak valid")
			return
		}
		created, err := h.service.Create(m)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, created)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}

func (h *LokasiHandler) HandleLokasiByID(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r, "/api/lokasi/")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id tidak valid")
		return
	}

	switch r.Method {
	case http.MethodGet:
		m, err := h.service.GetByID(id)
		if err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, m)

	case http.MethodPut:
		var m model.Lokasi
		if err := decodeJSON(r, &m); err != nil {
			writeError(w, http.StatusBadRequest, "data tidak valid")
			return
		}
		updated, err := h.service.Update(id, m)
		if err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, updated)

	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, map[string]string{"message": "lokasi berhasil dihapus"})

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}
