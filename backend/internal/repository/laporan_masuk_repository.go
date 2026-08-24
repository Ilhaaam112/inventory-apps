package repository

import "github.com/username/belajar_go/backend/internal/model"

// LaporanMasuk merinci setiap item barang masuk dalam satu periode.
// Kirim 0 pada lokasiID / supplierID untuk mengabaikan filter tersebut.
func (r *LaporanRepository) LaporanMasuk(start, end string, lokasiID, supplierID int) []model.LaporanMasukRow {
	rows, err := r.db.Query(`
		SELECT DATE_FORMAT(si.tanggal, '%Y-%m-%d'), si.code, sp.nama_supplier, l.nama_lokasi,
		       b.nama, s.nama_satuan, d.quantity, d.harga_beli
		FROM stock_in_details d
		JOIN stock_ins si     ON d.stock_in_id = si.id
		JOIN barang b         ON d.barang_id   = b.id
		JOIN lokasi l         ON si.lokasi_id  = l.id
		LEFT JOIN satuan s    ON b.satuan_id   = s.id
		LEFT JOIN supplier sp ON si.supplier_id = sp.id
		WHERE si.tanggal BETWEEN ? AND ?
		  AND (? = 0 OR si.lokasi_id = ?) AND (? = 0 OR si.supplier_id = ?)
		ORDER BY si.tanggal, si.id, d.id`,
		start, end, lokasiID, lokasiID, supplierID, supplierID)
	if err != nil {
		return []model.LaporanMasukRow{}
	}
	defer rows.Close()

	list := []model.LaporanMasukRow{}
	for rows.Next() {
		var x model.LaporanMasukRow
		if err := rows.Scan(&x.Tanggal, &x.Code, &x.NamaSupplier, &x.NamaLokasi,
			&x.NamaBarang, &x.NamaSatuan, &x.Quantity, &x.HargaBeli); err == nil {
			x.Subtotal = float64(x.Quantity) * x.HargaBeli
			list = append(list, x)
		}
	}
	return list
}
