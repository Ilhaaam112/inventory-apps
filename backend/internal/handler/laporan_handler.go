package handler

import (
	"net/http"

	"github.com/username/belajar_go/backend/internal/service"
)

type LaporanHandler struct {
	service *service.LaporanService
}

func NewLaporanHandler(s *service.LaporanService) *LaporanHandler {
	return &LaporanHandler{service: s}
}

// queryStr mengambil parameter teks dari URL, misal ?start=2026-08-01
func queryStr(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
