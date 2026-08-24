package repository

import "github.com/username/belajar_go/backend/internal/model"

// LaporanKeluar merinci setiap item barang keluar dalam satu periode.
// Kirim 0 pada lokasiID untuk mengabaikan filter gudang.
func (r *LaporanRepository) LaporanKeluar(start, end string, lokasiID int) []model.LaporanKeluarRow {
	rows, err := r.db.Query(`
		SELECT DATE_FORMAT(so.tanggal, '%Y-%m-%d'), so.code, l.nama_lokasi, so.tujuan,
		       b.nama, s.nama_satuan, d.quantity
		FROM stock_out_details d
		JOIN stock_outs so ON d.stock_out_id = so.id
		JOIN barang b      ON d.barang_id    = b.id
		JOIN lokasi l      ON so.lokasi_id   = l.id
		LEFT JOIN satuan s ON b.satuan_id    = s.id
		WHERE so.tanggal BETWEEN ? AND ?
		  AND (? = 0 OR so.lokasi_id = ?)
		ORDER BY so.tanggal, so.id, d.id`,
		start, end, lokasiID, lokasiID)
	if err != nil {
		return []model.LaporanKeluarRow{}
	}
	defer rows.Close()

	list := []model.LaporanKeluarRow{}
	for rows.Next() {
		var x model.LaporanKeluarRow
		if err := rows.Scan(&x.Tanggal, &x.Code, &x.NamaLokasi, &x.Tujuan,
			&x.NamaBarang, &x.NamaSatuan, &x.Quantity); err == nil {
			list = append(list, x)
		}
	}
	return list
}
