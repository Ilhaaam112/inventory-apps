package repository

import (
	"time"

	"github.com/username/belajar_go/backend/internal/model"
)

// Tren mengembalikan total unit masuk dan keluar per hari untuk rentang
// tanggal apa pun. Hari tanpa transaksi tetap dikembalikan dengan nilai 0,
// supaya grafik di frontend tidak bolong.
func (r *DashboardRepository) Tren(start, end string) []model.TrenHarian {
	hasil := map[string]model.TrenHarian{}

	rows, err := r.db.Query(`
		SELECT DATE_FORMAT(DATE(created_at), '%Y-%m-%d'),
		       COALESCE(SUM(CASE WHEN type IN ('IN','TRANSFER_IN')   THEN quantity ELSE 0 END), 0),
		       COALESCE(-SUM(CASE WHEN type IN ('OUT','TRANSFER_OUT') THEN quantity ELSE 0 END), 0)
		FROM stock_movements
		WHERE DATE(created_at) BETWEEN ? AND ?
		GROUP BY DATE(created_at)`, start, end)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var t model.TrenHarian
			if err := rows.Scan(&t.Tanggal, &t.Masuk, &t.Keluar); err == nil {
				hasil[t.Tanggal] = t
			}
		}
	}

	// Isi setiap hari dalam rentang, termasuk yang kosong.
	awal, err1 := time.Parse("2006-01-02", start)
	akhir, err2 := time.Parse("2006-01-02", end)
	if err1 != nil || err2 != nil {
		return []model.TrenHarian{}
	}

	tren := []model.TrenHarian{}
	for d := awal; !d.After(akhir); d = d.AddDate(0, 0, 1) {
		tgl := d.Format("2006-01-02")
		if t, ada := hasil[tgl]; ada {
			tren = append(tren, t)
		} else {
			tren = append(tren, model.TrenHarian{Tanggal: tgl})
		}
	}
	return tren
}
