package service

import (
	"time"

	"github.com/username/belajar_go/backend/internal/model"
)

// maksimalHariTren membatasi rentang supaya grafik tetap terbaca
// dan query tidak memindai terlalu banyak baris.
const maksimalHariTren = 92

func (s *DashboardService) Tren(start, end string) []model.TrenHarian {
	layout := "2006-01-02"
	now := time.Now()

	akhir, err := time.Parse(layout, end)
	if err != nil {
		akhir = now
	}
	awal, err := time.Parse(layout, start)
	if err != nil {
		awal = akhir.AddDate(0, 0, -6)
	}
	if awal.After(akhir) {
		awal, akhir = akhir, awal
	}
	// Potong dari sisi awal kalau rentangnya kelewat panjang.
	if akhir.Sub(awal).Hours()/24 > maksimalHariTren {
		awal = akhir.AddDate(0, 0, -maksimalHariTren)
	}

	return s.repo.Tren(awal.Format(layout), akhir.Format(layout))
}
