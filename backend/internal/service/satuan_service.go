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
	if err := idValid(id); err != nil {
		return model.Satuan{}, err
	}
	return s.repo.GetByID(id)
}

func (s *SatuanService) bersihkan(m model.Satuan) (model.Satuan, error) {
	nama, err := teksWajib(m.NamaSatuan, "nama satuan", 50)
	if err != nil {
		return model.Satuan{}, err
	}
	ket, err := teksOpsional(m.Keterangan, "keterangan", 255)
	if err != nil {
		return model.Satuan{}, err
	}
	m.NamaSatuan, m.Keterangan = nama, ket
	return m, nil
}

func (s *SatuanService) Create(m model.Satuan) (model.Satuan, error) {
	m, err := s.bersihkan(m)
	if err != nil {
		return model.Satuan{}, err
	}
	created, err := s.repo.Create(m)
	if err != nil {
		return model.Satuan{}, errors.New("gagal menyimpan satuan")
	}
	return created, nil
}

func (s *SatuanService) Update(id int, m model.Satuan) (model.Satuan, error) {
	if err := idValid(id); err != nil {
		return model.Satuan{}, err
	}
	m, err := s.bersihkan(m)
	if err != nil {
		return model.Satuan{}, err
	}
	return s.repo.Update(id, m)
}

func (s *SatuanService) Delete(id int) error {
	if err := idValid(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
