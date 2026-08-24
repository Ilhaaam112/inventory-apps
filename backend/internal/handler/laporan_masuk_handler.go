package handler

import "net/http"

// GET /api/laporan/barang-masuk?start=&end=&lokasi_id=&supplier_id=
func (h *LaporanHandler) HandleLaporanMasuk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	writeJSON(w, h.service.LaporanMasuk(
		queryStr(r, "start"), queryStr(r, "end"),
		queryInt(r, "lokasi_id"), queryInt(r, "supplier_id"),
	))
}
