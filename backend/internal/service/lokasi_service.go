package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type LokasiService struct {
	repo *repository.LokasiRepository
}

func NewLokasiService(repo *repository.LokasiRepository) *LokasiService {
	return &LokasiService{repo: repo}
}

func (s *LokasiService) GetAll() []model.Lokasi {
	return s.repo.GetAll()
}

func (s *LokasiService) GetByID(id int) (model.Lokasi, error) {
	if err := idValid(id); err != nil {
		return model.Lokasi{}, err
	}
	return s.repo.GetByID(id)
}

func (s *LokasiService) bersihkan(m model.Lokasi) (model.Lokasi, error) {
	nama, err := teksWajib(m.NamaLokasi, "nama lokasi", 100)
	if err != nil {
		return model.Lokasi{}, err
	}
	ket, err := teksOpsional(m.Keterangan, "keterangan", 255)
	if err != nil {
		return model.Lokasi{}, err
	}
	m.NamaLokasi, m.Keterangan = nama, ket
	return m, nil
}

func (s *LokasiService) Create(m model.Lokasi) (model.Lokasi, error) {
	m, err := s.bersihkan(m)
	if err != nil {
		return model.Lokasi{}, err
	}
	created, err := s.repo.Create(m)
	if err != nil {
		return model.Lokasi{}, errors.New("gagal menyimpan lokasi")
	}
	return created, nil
}

func (s *LokasiService) Update(id int, m model.Lokasi) (model.Lokasi, error) {
	if err := idValid(id); err != nil {
		return model.Lokasi{}, err
	}
	m, err := s.bersihkan(m)
	if err != nil {
		return model.Lokasi{}, err
	}
	return s.repo.Update(id, m)
}

func (s *LokasiService) Delete(id int) error {
	if err := idValid(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
