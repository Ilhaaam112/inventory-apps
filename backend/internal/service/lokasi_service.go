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
	return s.repo.GetByID(id)
}

func (s *LokasiService) Create(m model.Lokasi) (model.Lokasi, error) {
	if m.NamaLokasi == "" {
		return model.Lokasi{}, errors.New("nama lokasi wajib diisi")
	}
	return s.repo.Create(m), nil
}

func (s *LokasiService) Update(id int, m model.Lokasi) (model.Lokasi, error) {
	if m.NamaLokasi == "" {
		return model.Lokasi{}, errors.New("nama lokasi wajib diisi")
	}
	return s.repo.Update(id, m)
}

func (s *LokasiService) Delete(id int) error {
	return s.repo.Delete(id)
}
