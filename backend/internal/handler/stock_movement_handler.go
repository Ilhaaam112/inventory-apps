package handler

import (
	"net/http"

	"github.com/username/belajar_go/backend/internal/service"
)

type StockMovementHandler struct {
	service *service.StockMovementService
}

func NewStockMovementHandler(s *service.StockMovementService) *StockMovementHandler {
	return &StockMovementHandler{service: s}
}

// GET /api/stock-movements?lokasi_id=1&barang_id=2&limit=50
func (h *StockMovementHandler) HandleStockMovement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.GetAll(queryInt(r, "lokasi_id"), queryInt(r, "barang_id"), queryInt(r, "limit")))
}
