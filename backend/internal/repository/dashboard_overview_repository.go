package repository

import (
	"time"

	"github.com/username/belajar_go/backend/internal/model"
)

// Overview mengumpulkan angka ringkas dalam satu query,
// lalu melengkapinya dengan tren 7 hari terakhir.
func (r *DashboardRepository) Overview() model.DashboardOverview {
	var o model.DashboardOverview

	r.db.QueryRow(`
		SELECT
			(SELECT COUNT(*) FROM barang),
			(SELECT COUNT(*) FROM lokasi),
			(SELECT COUNT(*) FROM supplier),
			(SELECT COALESCE(SUM(quantity), 0) FROM warehouse_stocks),
			(SELECT COALESCE(SUM(ws.quantity * b.harga), 0)
			   FROM warehouse_stocks ws JOIN barang b ON ws.barang_id = b.id),
			(SELECT COUNT(*) FROM warehouse_stocks ws JOIN barang b ON ws.barang_id = b.id
			   WHERE ws.quantity > 0 AND b.stok_minimum > 0 AND ws.quantity <= b.stok_minimum),
			(SELECT COUNT(*) FROM warehouse_stocks WHERE quantity = 0),
			(SELECT COALESCE(SUM(quantity), 0) FROM stock_movements
			   WHERE type IN ('IN','TRANSFER_IN') AND DATE(created_at) = CURDATE()),
			(SELECT COALESCE(-SUM(quantity), 0) FROM stock_movements
			   WHERE type IN ('OUT','TRANSFER_OUT') AND DATE(created_at) = CURDATE()),
			(SELECT
				(SELECT COUNT(*) FROM stock_ins         WHERE tanggal BETWEEN ? AND ?) +
				(SELECT COUNT(*) FROM stock_outs        WHERE tanggal BETWEEN ? AND ?) +
				(SELECT COUNT(*) FROM stock_adjustments WHERE tanggal BETWEEN ? AND ?) +
				(SELECT COUNT(*) FROM stock_transfers   WHERE tanggal BETWEEN ? AND ?))`,
		awalBulanIni(), akhirBulanIni(), awalBulanIni(), akhirBulanIni(),
		awalBulanIni(), akhirBulanIni(), awalBulanIni(), akhirBulanIni(),
	).Scan(&o.TotalBarang, &o.TotalGudang, &o.TotalSupplier, &o.TotalUnit, &o.NilaiPersediaan,
		&o.StokMenipis, &o.StokHabis, &o.MasukHariIni, &o.KeluarHariIni, &o.TransaksiBulan)

	o.Tren = r.tren7Hari()
	return o
}

func awalBulanIni() string {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), 1, 0, 0, 0, 0, n.Location()).Format("2006-01-02")
}

func akhirBulanIni() string {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), 1, 0, 0, 0, 0, n.Location()).
		AddDate(0, 1, -1).Format("2006-01-02")
}

// tren7Hari selalu mengembalikan 7 baris; hari tanpa transaksi diisi nol,
// supaya grafik di frontend tidak bolong.
func (r *DashboardRepository) tren7Hari() []model.TrenHarian {
	hasil := map[string]model.TrenHarian{}

	rows, err := r.db.Query(`
		SELECT DATE_FORMAT(DATE(created_at), '%Y-%m-%d'),
		       COALESCE(SUM(CASE WHEN type IN ('IN','TRANSFER_IN')  THEN quantity ELSE 0 END), 0),
		       COALESCE(-SUM(CASE WHEN type IN ('OUT','TRANSFER_OUT') THEN quantity ELSE 0 END), 0)
		FROM stock_movements
		WHERE DATE(created_at) >= DATE_SUB(CURDATE(), INTERVAL 6 DAY)
		GROUP BY DATE(created_at)`)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var t model.TrenHarian
			if err := rows.Scan(&t.Tanggal, &t.Masuk, &t.Keluar); err == nil {
				hasil[t.Tanggal] = t
			}
		}
	}

	tren := make([]model.TrenHarian, 0, 7)
	for i := 6; i >= 0; i-- {
		tgl := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		if t, ada := hasil[tgl]; ada {
			tren = append(tren, t)
		} else {
			tren = append(tren, model.TrenHarian{Tanggal: tgl})
		}
	}
	return tren
}
