package service

import "github.com/username/belajar_go/backend/internal/model"

func (s *LaporanService) LaporanStok(lokasiID, kategoriID int) []model.LaporanStokRow {
	return s.repo.LaporanStok(lokasiID, kategoriID)
}
