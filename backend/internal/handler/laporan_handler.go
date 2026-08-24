package handler

import "github.com/username/belajar_go/backend/internal/service"

type LaporanHandler struct {
	service *service.LaporanService
}

func NewLaporanHandler(s *service.LaporanService) *LaporanHandler {
	return &LaporanHandler{service: s}
}
