package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type StockInService struct {
	repo *repository.StockInRepository
}

func NewStockInService(repo *repository.StockInRepository) *StockInService {
	return &StockInService{repo: repo}
}

func (s *StockInService) GetAll() []model.StockIn {
	return s.repo.GetAll()
}

func (s *StockInService) GetByID(id int) (model.StockIn, error) {
	return s.repo.GetByID(id)
}

func (s *StockInService) Create(in model.StockIn) (model.StockIn, error) {
	if in.LokasiID == 0 {
		return model.StockIn{}, errors.New("gudang tujuan wajib dipilih")
	}
	if in.Tanggal == "" {
		return model.StockIn{}, errors.New("tanggal wajib diisi")
	}
	if len(in.Details) == 0 {
		return model.StockIn{}, errors.New("detail barang minimal satu baris")
	}
	for _, d := range in.Details {
		if d.BarangID == 0 {
			return model.StockIn{}, errors.New("barang wajib dipilih pada setiap baris")
		}
		if d.Quantity <= 0 {
			return model.StockIn{}, errors.New("jumlah harus lebih dari 0")
		}
		if d.HargaBeli < 0 {
			return model.StockIn{}, errors.New("harga beli tidak boleh negatif")
		}
	}
	return s.repo.Create(in)
}
