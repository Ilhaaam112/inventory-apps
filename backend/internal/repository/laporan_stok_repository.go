package repository

import "github.com/username/belajar_go/backend/internal/model"

// LaporanStok menampilkan posisi stok saat ini.
// Kirim 0 pada lokasiID / kategoriID untuk mengabaikan filter tersebut.
func (r *LaporanRepository) LaporanStok(lokasiID, kategoriID int) []model.LaporanStokRow {
	rows, err := r.db.Query(`
		SELECT ws.lokasi_id, l.nama_lokasi, ws.barang_id, b.nama,
		       k.nama_kategori, s.nama_satuan, ws.quantity, b.harga
		FROM warehouse_stocks ws
		JOIN barang b        ON ws.barang_id  = b.id
		JOIN lokasi l        ON ws.lokasi_id  = l.id
		LEFT JOIN kategori k ON b.kategori_id = k.id
		LEFT JOIN satuan s   ON b.satuan_id   = s.id
		WHERE (? = 0 OR ws.lokasi_id = ?) AND (? = 0 OR b.kategori_id = ?)
		ORDER BY b.nama, l.nama_lokasi`,
		lokasiID, lokasiID, kategoriID, kategoriID)
	if err != nil {
		return []model.LaporanStokRow{}
	}
	defer rows.Close()

	list := []model.LaporanStokRow{}
	for rows.Next() {
		var x model.LaporanStokRow
		if err := rows.Scan(&x.LokasiID, &x.NamaLokasi, &x.BarangID, &x.NamaBarang,
			&x.NamaKategori, &x.NamaSatuan, &x.Quantity, &x.Harga); err == nil {
			x.Nilai = float64(x.Quantity) * x.Harga
			list = append(list, x)
		}
	}
	return list
}
