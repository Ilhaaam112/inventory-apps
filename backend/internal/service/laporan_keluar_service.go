package service

import "github.com/username/belajar_go/backend/internal/model"

func (s *LaporanService) LaporanKeluar(start, end string, lokasiID int) []model.LaporanKeluarRow {
	start, end = periode(start, end)
	return s.repo.LaporanKeluar(start, end, lokasiID)
}
