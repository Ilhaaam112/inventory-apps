package repository

import (
	"log"
	"time"

	"github.com/username/belajar_go/backend/internal/model"
)

// Overview mengumpulkan angka ringkas dalam satu query,
// lalu melengkapinya dengan tren pada rentang yang dipilih user.
func (r *DashboardRepository) Overview(start, end string) model.DashboardOverview {
	var o model.DashboardOverview
	o.Start, o.End = start, end

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

	o.Tren = r.tren(start, end)
	for _, t := range o.Tren {
		o.MasukPeriode += t.Masuk
		o.KeluarPeriode += t.Keluar
	}
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

// tren mengembalikan satu baris untuk SETIAP hari dalam rentang,
// termasuk hari tanpa transaksi (diisi nol) supaya grafik tidak bolong.
// Rentang dibatasi 92 hari agar payload tidak membengkak.
func (r *DashboardRepository) tren(start, end string) []model.TrenHarian {
	hasil := map[string]model.TrenHarian{}

	// GROUP BY memakai alias `tgl`, bukan DATE(created_at). Kalau SELECT
	// dan GROUP BY memakai ekspresi yang berbeda, MySQL 8 dengan
	// sql_mode=only_full_group_by menolaknya dengan error 1055.
	//
	// Rentang dibandingkan langsung ke created_at (bukan DATE(created_at))
	// supaya index idx_mv_created tetap terpakai.
	rows, err := r.db.Query(`
        SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS tgl,
               COALESCE(SUM(CASE WHEN type IN ('IN','TRANSFER_IN')   THEN quantity ELSE 0 END), 0) AS masuk,
               COALESCE(-SUM(CASE WHEN type IN ('OUT','TRANSFER_OUT') THEN quantity ELSE 0 END), 0) AS keluar
        FROM stock_movements
        WHERE created_at >= ? AND created_at < DATE_ADD(?, INTERVAL 1 DAY)
        GROUP BY tgl`, start, end)

	if err != nil {
		// Jangan menelan error query diam-diam: tanpa log ini, grafik
		// kosong karena error terlihat sama persis dengan grafik kosong
		// karena memang belum ada data.
		log.Println("query tren gagal:", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var t model.TrenHarian
			if err := rows.Scan(&t.Tanggal, &t.Masuk, &t.Keluar); err != nil {
				log.Println("scan tren gagal:", err)
				continue
			}
			hasil[t.Tanggal] = t
		}
	}

	awal, err1 := time.Parse("2006-01-02", start)
	akhir, err2 := time.Parse("2006-01-02", end)
	if err1 != nil || err2 != nil {
		return []model.TrenHarian{}
	}

	tren := []model.TrenHarian{}
	for d := awal; !d.After(akhir) && len(tren) < 92; d = d.AddDate(0, 0, 1) {
		tgl := d.Format("2006-01-02")
		if t, ada := hasil[tgl]; ada {
			tren = append(tren, t)
		} else {
			tren = append(tren, model.TrenHarian{Tanggal: tgl})
		}
	}
	return tren
}
