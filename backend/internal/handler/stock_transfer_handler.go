package handler

import (
	"net/http"

	"github.com/username/belajar_go/backend/internal/middleware"
	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/service"
)

type StockTransferHandler struct {
	service *service.StockTransferService
}

func NewStockTransferHandler(s *service.StockTransferService) *StockTransferHandler {
	return &StockTransferHandler{service: s}
}

func (h *StockTransferHandler) HandleTransfer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, h.service.GetAll())

	case http.MethodPost:
		var t model.StockTransfer
		if err := decodeJSON(r, &t); err != nil {
			writeError(w, http.StatusBadRequest, "data tidak valid")
			return
		}
		if uid := middleware.UserID(r); uid > 0 {
			t.UserID = &uid
		}
		created, err := h.service.Create(t)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, created)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
	}
}

func (h *StockTransferHandler) HandleTransferByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	id, err := idFromPath(r, "/api/stock-transfer/")
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
