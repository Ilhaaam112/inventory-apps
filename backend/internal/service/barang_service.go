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
	if err := idValid(id); err != nil {
		return model.Barang{}, err
	}
	return s.repo.GetByID(id)
}

func (s *BarangService) bersihkan(b model.Barang) (model.Barang, error) {
	nama, err := teksWajib(b.Nama, "nama barang", 150)
	if err != nil {
		return model.Barang{}, err
	}
	if b.Harga < 0 {
		return model.Barang{}, errors.New("harga tidak boleh negatif")
	}
	if b.Harga > 999999999999 {
		return model.Barang{}, errors.New("harga terlalu besar")
	}
	if b.StokMinimum < 0 {
		return model.Barang{}, errors.New("stok minimum tidak boleh negatif")
	}
	if b.SatuanID == nil || *b.SatuanID <= 0 {
		return model.Barang{}, errors.New("satuan wajib dipilih")
	}
	if b.KategoriID != nil && *b.KategoriID <= 0 {
		b.KategoriID = nil
	}
	b.Nama = nama
	return b, nil
}

func (s *BarangService) Create(b model.Barang) (model.Barang, error) {
	b, err := s.bersihkan(b)
	if err != nil {
		return model.Barang{}, err
	}
	created, err := s.repo.Create(b)
	if err != nil {
		return model.Barang{}, errors.New("gagal menyimpan barang")
	}
	return created, nil
}

func (s *BarangService) Update(id int, b model.Barang) (model.Barang, error) {
	if err := idValid(id); err != nil {
		return model.Barang{}, err
	}
	b, err := s.bersihkan(b)
	if err != nil {
		return model.Barang{}, err
	}
	return s.repo.Update(id, b)
}

func (s *BarangService) Delete(id int) error {
	if err := idValid(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
