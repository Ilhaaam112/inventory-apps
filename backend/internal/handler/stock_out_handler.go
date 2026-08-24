package handler

import (
	"encoding/json"
	"net/http"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type StockOutHandler struct {
	service *service.StockOutService
}

func NewStockOutHandler(s *service.StockOutService) *StockOutHandler {
	return &StockOutHandler{service: s}
}

func (h *StockOutHandler) HandleStockOut(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, h.service.GetAll())

	case http.MethodPost:
		var out model.StockOut
		if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
			writeError(w, http.StatusBadRequest, "data tidak valid")
			return
		}
		created, err := h.service.Create(out)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, created)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}

func (h *StockOutHandler) HandleStockOutByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	id, err := idFromPath(r, "/api/stock-out/")
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
