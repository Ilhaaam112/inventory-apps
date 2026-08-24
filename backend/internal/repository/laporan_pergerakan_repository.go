package repository

import "github.com/username/belajar_go/backend/internal/model"

// LaporanPergerakan merekap mutasi stok per barang per gudang.
//
// Saldo akhir dihitung mundur dari stok sekarang dikurangi pergerakan
// SETELAH periode, lalu saldo awal = saldo akhir dikurangi total mutasi
// dalam periode. Cara ini tetap benar walaupun stok awal dulu diisi
// langsung lewat migrasi tanpa melalui stock_movements.
func (r *LaporanRepository) LaporanPergerakan(start, end string, lokasiID int) []model.LaporanPergerakanRow {
	rows, err := r.db.Query(`
		SELECT ws.lokasi_id, l.nama_lokasi, ws.barang_id, b.nama, s.nama_satuan, ws.quantity,
		       COALESCE((SELECT SUM(m.quantity) FROM stock_movements m
		                 WHERE m.lokasi_id = ws.lokasi_id AND m.barang_id = ws.barang_id
		                   AND DATE(m.created_at) > ?), 0) AS setelah,
		       COALESCE((SELECT SUM(m.quantity) FROM stock_movements m
		                 WHERE m.lokasi_id = ws.lokasi_id AND m.barang_id = ws.barang_id
		                   AND DATE(m.created_at) BETWEEN ? AND ?), 0) AS mutasi,
		       COALESCE((SELECT SUM(m.quantity) FROM stock_movements m
		                 WHERE m.lokasi_id = ws.lokasi_id AND m.barang_id = ws.barang_id
		                   AND m.type IN ('IN','TRANSFER_IN')
		                   AND DATE(m.created_at) BETWEEN ? AND ?), 0) AS masuk,
		       COALESCE((SELECT SUM(m.quantity) FROM stock_movements m
		                 WHERE m.lokasi_id = ws.lokasi_id AND m.barang_id = ws.barang_id
		                   AND m.type IN ('OUT','TRANSFER_OUT')
		                   AND DATE(m.created_at) BETWEEN ? AND ?), 0) AS keluar,
		       COALESCE((SELECT SUM(m.quantity) FROM stock_movements m
		                 WHERE m.lokasi_id = ws.lokasi_id AND m.barang_id = ws.barang_id
		                   AND m.type = 'ADJUSTMENT'
		                   AND DATE(m.created_at) BETWEEN ? AND ?), 0) AS penyesuaian
		FROM warehouse_stocks ws
		JOIN barang b      ON ws.barang_id = b.id
		JOIN lokasi l      ON ws.lokasi_id = l.id
		LEFT JOIN satuan s ON b.satuan_id  = s.id
		WHERE (? = 0 OR ws.lokasi_id = ?)
		ORDER BY b.nama, l.nama_lokasi`,
		end, start, end, start, end, start, end, start, end, lokasiID, lokasiID)
	if err != nil {
		return []model.LaporanPergerakanRow{}
	}
	defer rows.Close()

	list := []model.LaporanPergerakanRow{}
	for rows.Next() {
		var x model.LaporanPergerakanRow
		var stokSekarang, setelah, mutasi, keluar int
		if err := rows.Scan(&x.LokasiID, &x.NamaLokasi, &x.BarangID, &x.NamaBarang, &x.NamaSatuan,
			&stokSekarang, &setelah, &mutasi, &x.Masuk, &keluar, &x.Penyesuaian); err != nil {
			continue
		}
		x.SaldoAkhir = stokSekarang - setelah
		x.SaldoAwal = x.SaldoAkhir - mutasi
		x.Keluar = -keluar // disimpan negatif di stock_movements
		list = append(list, x)
	}
	return list
}
