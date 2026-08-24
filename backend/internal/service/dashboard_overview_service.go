package service

import (
	"time"

	"github.com/username/belajar_go/backend/internal/model"
)

// periodeDashboard mengisi tanggal kosong dengan default 7 hari terakhir,
// menukar tanggal yang terbalik, dan membatasi rentang maksimal 92 hari.
func periodeDashboard(start, end string) (string, string) {
	const f = "2006-01-02"
	now := time.Now()

	if end == "" {
		end = now.Format(f)
	}
	if start == "" {
		start = now.AddDate(0, 0, -6).Format(f)
	}
	if start > end {
		start, end = end, start
	}

	awal, err1 := time.Parse(f, start)
	akhir, err2 := time.Parse(f, end)
	if err1 != nil || err2 != nil {
		return now.AddDate(0, 0, -6).Format(f), now.Format(f)
	}
	if akhir.Sub(awal).Hours() > 92*24 {
		start = akhir.AddDate(0, 0, -91).Format(f)
	}
	return start, end
}

func (s *DashboardService) Overview(start, end string) model.DashboardOverview {
	start, end = periodeDashboard(start, end)
	return s.repo.Overview(start, end)
}
