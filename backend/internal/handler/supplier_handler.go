package handler

import (
	"net/http"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type SupplierHandler struct {
	service *service.SupplierService
}

func NewSupplierHandler(s *service.SupplierService) *SupplierHandler {
	return &SupplierHandler{service: s}
}

func (h *SupplierHandler) HandleSupplier(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, h.service.GetAll())

	case http.MethodPost:
		var m model.Supplier
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

func (h *SupplierHandler) HandleSupplierByID(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r, "/api/supplier/")
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
		var m model.Supplier
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
		writeJSON(w, map[string]string{"message": "supplier berhasil dihapus"})

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}
