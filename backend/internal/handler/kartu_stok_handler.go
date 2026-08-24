package handler

import "net/http"

// GET /api/laporan/kartu-stok?barang_id=&lokasi_id=&start=&end=
func (h *LaporanHandler) HandleKartuStok(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method tidak diizinkan")
		return
	}
	data, err := h.service.KartuStok(
		queryInt(r, "barang_id"), queryInt(r, "lokasi_id"),
		queryStr(r, "start"), queryStr(r, "end"),
	)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, data)
}
