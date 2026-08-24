package handler

import "net/http"

// GET /api/laporan/stok?lokasi_id=&kategori_id=
func (h *LaporanHandler) HandleLaporanStok(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.LaporanStok(queryInt(r, "lokasi_id"), queryInt(r, "kategori_id")))
}
