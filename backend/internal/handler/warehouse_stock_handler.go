package handler

import (
	"net/http"

	"github.com/username/belajar_go/backend/internal/service"
)

type WarehouseStockHandler struct {
	service *service.WarehouseStockService
}

func NewWarehouseStockHandler(s *service.WarehouseStockService) *WarehouseStockHandler {
	return &WarehouseStockHandler{service: s}
}

// GET /api/warehouse-stocks?lokasi_id=1&barang_id=2
func (h *WarehouseStockHandler) HandleWarehouseStock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.GetAll(queryInt(r, "lokasi_id"), queryInt(r, "barang_id")))
}
