package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type SatuanService struct {
	repo *repository.SatuanRepository
}

func NewSatuanService(repo *repository.SatuanRepository) *SatuanService {
	return &SatuanService{repo: repo}
}

func (s *SatuanService) GetAll() []model.Satuan {
	return s.repo.GetAll()
}

func (s *SatuanService) GetByID(id int) (model.Satuan, error) {
	return s.repo.GetByID(id)
}

func (s *SatuanService) Create(m model.Satuan) (model.Satuan, error) {
	if m.NamaSatuan == "" {
		return model.Satuan{}, errors.New("nama satuan wajib diisi")
	}
	return s.repo.Create(m), nil
}

func (s *SatuanService) Update(id int, m model.Satuan) (model.Satuan, error) {
	if m.NamaSatuan == "" {
		return model.Satuan{}, errors.New("nama satuan wajib diisi")
	}
	return s.repo.Update(id, m)
}

func (s *SatuanService) Delete(id int) error {
	return s.repo.Delete(id)
}
