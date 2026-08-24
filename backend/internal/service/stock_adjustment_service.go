package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type StockAdjustmentService struct {
	repo *repository.StockAdjustmentRepository
}

func NewStockAdjustmentService(repo *repository.StockAdjustmentRepository) *StockAdjustmentService {
	return &StockAdjustmentService{repo: repo}
}

func (s *StockAdjustmentService) GetAll() []model.StockAdjustment {
	return s.repo.GetAll()
}

func (s *StockAdjustmentService) GetByID(id int) (model.StockAdjustment, error) {
	return s.repo.GetByID(id)
}

func (s *StockAdjustmentService) Create(a model.StockAdjustment) (model.StockAdjustment, error) {
	if a.LokasiID == 0 {
		return model.StockAdjustment{}, errors.New("gudang wajib dipilih")
	}
	if a.Tanggal == "" {
		return model.StockAdjustment{}, errors.New("tanggal wajib diisi")
	}
	if a.Alasan == "" {
		return model.StockAdjustment{}, errors.New("alasan penyesuaian wajib diisi")
	}
	if len(a.Details) == 0 {
		return model.StockAdjustment{}, errors.New("detail barang minimal satu baris")
	}
	for _, d := range a.Details {
		if d.BarangID == 0 {
			return model.StockAdjustment{}, errors.New("barang wajib dipilih pada setiap baris")
		}
		if d.ActualStock < 0 {
			return model.StockAdjustment{}, errors.New("stok fisik tidak boleh negatif")
		}
	}
	return s.repo.Create(a)
}
