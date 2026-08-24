package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type StockTransferService struct {
	repo *repository.StockTransferRepository
}

func NewStockTransferService(repo *repository.StockTransferRepository) *StockTransferService {
	return &StockTransferService{repo: repo}
}

func (s *StockTransferService) GetAll() []model.StockTransfer {
	return s.repo.GetAll()
}

func (s *StockTransferService) GetByID(id int) (model.StockTransfer, error) {
	return s.repo.GetByID(id)
}

func (s *StockTransferService) Create(t model.StockTransfer) (model.StockTransfer, error) {
	if t.FromLokasiID == 0 || t.ToLokasiID == 0 {
		return model.StockTransfer{}, errors.New("gudang asal dan tujuan wajib dipilih")
	}
	if t.FromLokasiID == t.ToLokasiID {
		return model.StockTransfer{}, errors.New("gudang asal dan tujuan tidak boleh sama")
	}
	if t.Tanggal == "" {
		return model.StockTransfer{}, errors.New("tanggal wajib diisi")
	}
	if len(t.Details) == 0 {
		return model.StockTransfer{}, errors.New("detail barang minimal satu baris")
	}
	for _, d := range t.Details {
		if d.BarangID == 0 {
			return model.StockTransfer{}, errors.New("barang wajib dipilih pada setiap baris")
		}
		if d.Quantity <= 0 {
			return model.StockTransfer{}, errors.New("jumlah harus lebih dari 0")
		}
	}
	return s.repo.Create(t)
}
