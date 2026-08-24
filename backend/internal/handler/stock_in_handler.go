package handler

import (
	"encoding/json"
	"net/http"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type StockInHandler struct {
	service *service.StockInService
}

func NewStockInHandler(s *service.StockInService) *StockInHandler {
	return &StockInHandler{service: s}
}

func (h *StockInHandler) HandleStockIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, h.service.GetAll())

	case http.MethodPost:
		var in model.StockIn
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			writeError(w, http.StatusBadRequest, "data tidak valid")
			return
		}
		created, err := h.service.Create(in)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, created)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}

func (h *StockInHandler) HandleStockInByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	id, err := idFromPath(r, "/api/stock-in/")
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
