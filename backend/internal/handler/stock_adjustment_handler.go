package handler

import (
	"encoding/json"
	"net/http"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type StockAdjustmentHandler struct {
	service *service.StockAdjustmentService
}

func NewStockAdjustmentHandler(s *service.StockAdjustmentService) *StockAdjustmentHandler {
	return &StockAdjustmentHandler{service: s}
}

func (h *StockAdjustmentHandler) HandleAdjustment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, h.service.GetAll())

	case http.MethodPost:
		var a model.StockAdjustment
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			writeError(w, http.StatusBadRequest, "data tidak valid")
			return
		}
		created, err := h.service.Create(a)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, created)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}

func (h *StockAdjustmentHandler) HandleAdjustmentByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	id, err := idFromPath(r, "/api/stock-adjustment/")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id tidak valid")
		return
	}
	data, err := h.service.GetByID(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, data)
}
