package handler

import "net/http"

// GET /api/dashboard/stok-menipis?lokasi_id=&minimum=&limit=
func (h *DashboardHandler) HandleStokMenipis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.StokMenipis(
		queryInt(r, "lokasi_id"), queryInt(r, "minimum"), queryInt(r, "limit"),
	))
}
