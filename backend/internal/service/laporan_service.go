package service

import (
	"time"

	"github.com/username/belajar_go/backend/internal/repository"
)

type LaporanService struct {
	repo *repository.LaporanRepository
}

func NewLaporanService(repo *repository.LaporanRepository) *LaporanService {
	return &LaporanService{repo: repo}
}

// periode mengisi tanggal kosong dengan default: awal bulan ini sampai hari ini.
// Kalau terbalik (start > end), keduanya ditukar supaya query tetap masuk akal.
func periode(start, end string) (string, string) {
	now := time.Now()
	if start == "" {
		start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
	}
	if end == "" {
		end = now.Format("2006-01-02")
	}
	if start > end {
		start, end = end, start
	}
	return start, end
}
