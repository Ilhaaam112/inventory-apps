package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type StockOutService struct {
	repo *repository.StockOutRepository
}

func NewStockOutService(repo *repository.StockOutRepository) *StockOutService {
	return &StockOutService{repo: repo}
}

func (s *StockOutService) GetAll() []model.StockOut {
	return s.repo.GetAll()
}

func (s *StockOutService) GetByID(id int) (model.StockOut, error) {
	return s.repo.GetByID(id)
}

func (s *StockOutService) Create(out model.StockOut) (model.StockOut, error) {
	if out.LokasiID == 0 {
		return model.StockOut{}, errors.New("gudang asal wajib dipilih")
	}
	if out.Tanggal == "" {
		return model.StockOut{}, errors.New("tanggal wajib diisi")
	}
	if len(out.Details) == 0 {
		return model.StockOut{}, errors.New("detail barang minimal satu baris")
	}
	for _, d := range out.Details {
		if d.BarangID == 0 {
			return model.StockOut{}, errors.New("barang wajib dipilih pada setiap baris")
		}
		if d.Quantity <= 0 {
			return model.StockOut{}, errors.New("jumlah harus lebih dari 0")
		}
	}
	// Kecukupan stok tidak dicek di sini, tetapi di dalam database transaction,
	// supaya angkanya tidak keburu berubah sebelum transaksi disimpan.
	return s.repo.Create(out)
}
