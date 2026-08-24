package service

import (
	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type WarehouseStockService struct {
	repo *repository.WarehouseStockRepository
}

func NewWarehouseStockService(repo *repository.WarehouseStockRepository) *WarehouseStockService {
	return &WarehouseStockService{repo: repo}
}

func (s *WarehouseStockService) GetAll(lokasiID, barangID int) []model.WarehouseStock {
	return s.repo.GetAll(lokasiID, barangID)
}
