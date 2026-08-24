package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type BarangService struct {
	repo *repository.BarangRepository
}

func NewBarangService(repo *repository.BarangRepository) *BarangService {
	return &BarangService{repo: repo}
}

func (s *BarangService) GetAll() []model.Barang {
	return s.repo.GetAll()
}

func (s *BarangService) GetByID(id int) (model.Barang, error) {
	return s.repo.GetByID(id)
}

func (s *BarangService) Create(b model.Barang) (model.Barang, error) {
	if b.Nama == "" {
		return model.Barang{}, errors.New("nama barang wajib diisi")
	}
	if b.Harga < 0 {
		return model.Barang{}, errors.New("harga tidak boleh negatif")
	}
	if b.SatuanID == nil {
		return model.Barang{}, errors.New("satuan wajib dipilih")
	}
	return s.repo.Create(b), nil
}

func (s *BarangService) Update(id int, b model.Barang) (model.Barang, error) {
	if b.Nama == "" {
		return model.Barang{}, errors.New("nama barang wajib diisi")
	}
	if b.SatuanID == nil {
		return model.Barang{}, errors.New("satuan wajib dipilih")
	}
	return s.repo.Update(id, b)
}

func (s *BarangService) Delete(id int) error {
	return s.repo.Delete(id)
}