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
	if err := idValid(id); err != nil {
		return model.Supplier{}, err
	}
	return s.repo.GetByID(id)
}

func (s *SupplierService) bersihkan(m model.Supplier) (model.Supplier, error) {
	nama, err := teksWajib(m.NamaSupplier, "nama supplier", 150)
	if err != nil {
		return model.Supplier{}, err
	}
	kontak, err := teksOpsional(m.Kontak, "kontak", 100)
	if err != nil {
		return model.Supplier{}, err
	}
	alamat, err := teksOpsional(m.Alamat, "alamat", 255)
	if err != nil {
		return model.Supplier{}, err
	}
	m.NamaSupplier, m.Kontak, m.Alamat = nama, kontak, alamat
	return m, nil
}

func (s *SupplierService) Create(m model.Supplier) (model.Supplier, error) {
	m, err := s.bersihkan(m)
	if err != nil {
		return model.Supplier{}, err
	}
	created, err := s.repo.Create(m)
	if err != nil {
		return model.Supplier{}, errors.New("gagal menyimpan supplier")
	}
	return created, nil
}

func (s *SupplierService) Update(id int, m model.Supplier) (model.Supplier, error) {
	if err := idValid(id); err != nil {
		return model.Supplier{}, err
	}
	m, err := s.bersihkan(m)
	if err != nil {
		return model.Supplier{}, err
	}
	return s.repo.Update(id, m)
}

func (s *SupplierService) Delete(id int) error {
	if err := idValid(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
