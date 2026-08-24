package handler

import "net/http"

// GET /api/dashboard/overview?start=2026-08-01&end=2026-08-24
// Tanpa parameter, service memakai rentang bawaan (7 hari terakhir).
func (h *DashboardHandler) HandleOverview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.Overview(queryStr(r, "start"), queryStr(r, "end")))
}
