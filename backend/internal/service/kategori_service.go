package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type KategoriService struct {
	repo *repository.KategoriRepository
}

func NewKategoriService(repo *repository.KategoriRepository) *KategoriService {
	return &KategoriService{repo: repo}
}

func (s *KategoriService) GetAll() []model.Kategori {
	return s.repo.GetAll()
}

func (s *KategoriService) GetByID(id int) (model.Kategori, error) {
	return s.repo.GetByID(id)
}

func (s *KategoriService) Create(k model.Kategori) (model.Kategori, error) {
	if k.NamaKategori == "" {
		return model.Kategori{}, errors.New("nama kategori wajib diisi")
	}
	return s.repo.Create(k), nil
}

func (s *KategoriService) Update(id int, k model.Kategori) (model.Kategori, error) {
	if k.NamaKategori == "" {
		return model.Kategori{}, errors.New("nama kategori wajib diisi")
	}
	return s.repo.Update(id, k)
}

func (s *KategoriService) Delete(id int) error {
	return s.repo.Delete(id)
}