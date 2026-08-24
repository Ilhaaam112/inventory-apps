package service

import (
	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type StockMovementService struct {
	repo *repository.StockMovementRepository
}

func NewStockMovementService(repo *repository.StockMovementRepository) *StockMovementService {
	return &StockMovementService{repo: repo}
}

func (s *StockMovementService) GetAll(lokasiID, barangID, limit int) []model.StockMovement {
	return s.repo.GetAll(lokasiID, barangID, limit)
}
