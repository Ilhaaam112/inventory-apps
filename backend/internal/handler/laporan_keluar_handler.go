package handler

import "net/http"

// GET /api/laporan/barang-keluar?start=&end=&lokasi_id=
func (h *LaporanHandler) HandleLaporanKeluar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.LaporanKeluar(
		queryStr(r, "start"), queryStr(r, "end"), queryInt(r, "lokasi_id"),
	))
}
