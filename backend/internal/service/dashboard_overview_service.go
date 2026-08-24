package service

import "github.com/username/belajar_go/backend/internal/model"

func (s *DashboardService) Overview() model.DashboardOverview {
	return s.repo.Overview()
}
