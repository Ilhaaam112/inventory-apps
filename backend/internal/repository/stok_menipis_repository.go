package repository

import "github.com/username/belajar_go/backend/internal/model"

// StokMenipis mencari barang yang stoknya sudah menyentuh ambang batas.
// Barang dengan stok_minimum 0 hanya muncul kalau stoknya benar-benar habis,
// kecuali parameter minimum diisi sebagai ambang batas cadangan.
func (r *DashboardRepository) StokMenipis(lokasiID, minimum, limit int) []model.StokMenipisRow {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	rows, err := r.db.Query(`
		SELECT ws.lokasi_id, l.nama_lokasi, ws.barang_id, b.nama, s.nama_satuan,
		       ws.quantity, b.stok_minimum
		FROM warehouse_stocks ws
		JOIN barang b      ON ws.barang_id = b.id
		JOIN lokasi l      ON ws.lokasi_id = l.id
		LEFT JOIN satuan s ON b.satuan_id  = s.id
		WHERE (? = 0 OR ws.lokasi_id = ?)
		  AND ws.quantity <= IF(b.stok_minimum > 0, b.stok_minimum, ?)
		ORDER BY (ws.quantity - b.stok_minimum) ASC, b.nama
		LIMIT ?`,
		lokasiID, lokasiID, minimum, limit)
	if err != nil {
		return []model.StokMenipisRow{}
	}
	defer rows.Close()

	list := []model.StokMenipisRow{}
	for rows.Next() {
		var x model.StokMenipisRow
		if err := rows.Scan(&x.LokasiID, &x.NamaLokasi, &x.BarangID, &x.NamaBarang,
			&x.NamaSatuan, &x.Quantity, &x.StokMinimum); err != nil {
			continue
		}
		x.Habis = x.Quantity == 0
		if x.StokMinimum > x.Quantity {
			x.Kekurangan = x.StokMinimum - x.Quantity
		}
		list = append(list, x)
	}
	return list
}
