package service

import "github.com/username/belajar_go/backend/internal/model"

func (s *LaporanService) LaporanMasuk(start, end string, lokasiID, supplierID int) []model.LaporanMasukRow {
	start, end = periode(start, end)
	return s.repo.LaporanMasuk(start, end, lokasiID, supplierID)
}
