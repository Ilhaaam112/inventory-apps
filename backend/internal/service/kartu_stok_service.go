package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

func (s *LaporanService) KartuStok(barangID, lokasiID int, start, end string) (model.KartuStok, error) {
	if barangID == 0 {
		return model.KartuStok{}, errors.New("barang wajib dipilih")
	}
	if lokasiID == 0 {
		return model.KartuStok{}, errors.New("gudang wajib dipilih")
	}
	start, end = periode(start, end)
	return s.repo.KartuStok(barangID, lokasiID, start, end)
}
