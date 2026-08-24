package handler

import "github.com/username/belajar_go/backend/internal/service"

type DashboardHandler struct {
	service *service.DashboardService
}

func NewDashboardHandler(s *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: s}
}
