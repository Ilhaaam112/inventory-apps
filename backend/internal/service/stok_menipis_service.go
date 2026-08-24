package service

import "github.com/username/belajar_go/backend/internal/model"

func (s *DashboardService) StokMenipis(lokasiID, minimum, limit int) []model.StokMenipisRow {
	return s.repo.StokMenipis(lokasiID, minimum, limit)
}
