package service

import "github.com/username/belajar_go/backend/internal/model"

func (s *DashboardService) Aktivitas(limit int) []model.AktivitasRow {
	return s.repo.Aktivitas(limit)
}
