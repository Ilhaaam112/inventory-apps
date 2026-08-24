package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type SupplierService struct {
	repo *repository.SupplierRepository
}

func NewSupplierService(repo *repository.SupplierRepository) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s *SupplierService) GetAll() []model.Supplier {
	return s.repo.GetAll()
}

func (s *SupplierService) GetByID(id int) (model.Supplier, error) {
	return s.repo.GetByID(id)
}

func (s *SupplierService) Create(m model.Supplier) (model.Supplier, error) {
	if m.NamaSupplier == "" {
		return model.Supplier{}, errors.New("nama supplier wajib diisi")
	}
	return s.repo.Create(m), nil
}

func (s *SupplierService) Update(id int, m model.Supplier) (model.Supplier, error) {
	if m.NamaSupplier == "" {
		return model.Supplier{}, errors.New("nama supplier wajib diisi")
	}
	return s.repo.Update(id, m)
}

func (s *SupplierService) Delete(id int) error {
	return s.repo.Delete(id)
}
