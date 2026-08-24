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
	if err := idValid(id); err != nil {
		return model.Kategori{}, err
	}
	return s.repo.GetByID(id)
}

func (s *KategoriService) bersihkan(k model.Kategori) (model.Kategori, error) {
	nama, err := teksWajib(k.NamaKategori, "nama kategori", 100)
	if err != nil {
		return model.Kategori{}, err
	}
	k.NamaKategori = nama
	return k, nil
}

func (s *KategoriService) Create(k model.Kategori) (model.Kategori, error) {
	k, err := s.bersihkan(k)
	if err != nil {
		return model.Kategori{}, err
	}
	created, err := s.repo.Create(k)
	if err != nil {
		return model.Kategori{}, errors.New("gagal menyimpan kategori")
	}
	return created, nil
}

func (s *KategoriService) Update(id int, k model.Kategori) (model.Kategori, error) {
	if err := idValid(id); err != nil {
		return model.Kategori{}, err
	}
	k, err := s.bersihkan(k)
	if err != nil {
		return model.Kategori{}, err
	}
	return s.repo.Update(id, k)
}

func (s *KategoriService) Delete(id int) error {
	if err := idValid(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
