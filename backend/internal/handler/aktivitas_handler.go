package handler

import "net/http"

// GET /api/dashboard/aktivitas?limit=
func (h *DashboardHandler) HandleAktivitas(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.Aktivitas(queryInt(r, "limit")))
}
