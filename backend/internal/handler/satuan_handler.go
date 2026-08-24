package handler

import (
	"net/http"

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
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, h.service.GetAll())

	case http.MethodPost:
		var m model.Satuan
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

func (h *SatuanHandler) HandleSatuanByID(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r, "/api/satuan/")
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
		var m model.Satuan
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
		writeJSON(w, map[string]string{"message": "satuan berhasil dihapus"})

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}
