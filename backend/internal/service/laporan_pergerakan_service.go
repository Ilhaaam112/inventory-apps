package service

import "github.com/username/belajar_go/backend/internal/model"

func (s *LaporanService) LaporanPergerakan(start, end string, lokasiID int) []model.LaporanPergerakanRow {
	start, end = periode(start, end)
	return s.repo.LaporanPergerakan(start, end, lokasiID)
}
