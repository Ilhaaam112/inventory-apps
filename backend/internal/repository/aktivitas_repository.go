package repository

import "github.com/username/belajar_go/backend/internal/model"

// Aktivitas mengambil pergerakan stok terbaru beserta nomor transaksi
// dan nama petugas yang menyimpannya.
func (r *DashboardRepository) Aktivitas(limit int) []model.AktivitasRow {
	if limit <= 0 || limit > 200 {
		limit = 20
	}
	rows, err := r.db.Query(`
		SELECT DATE_FORMAT(m.created_at, '%Y-%m-%d %H:%i'), m.type,
		       COALESCE(
		           CASE m.reference_type
		               WHEN 'stock_in'         THEN (SELECT code FROM stock_ins         WHERE id = m.reference_id)
		               WHEN 'stock_out'        THEN (SELECT code FROM stock_outs        WHERE id = m.reference_id)
		               WHEN 'stock_adjustment' THEN (SELECT code FROM stock_adjustments WHERE id = m.reference_id)
		               WHEN 'stock_transfer'   THEN (SELECT code FROM stock_transfers   WHERE id = m.reference_id)
		           END, m.reference_type),
		       b.nama, l.nama_lokasi, m.quantity, m.stock_after,
		       (SELECT u.nama_lengkap FROM users u WHERE u.id =
		           CASE m.reference_type
		               WHEN 'stock_in'         THEN (SELECT user_id FROM stock_ins         WHERE id = m.reference_id)
		               WHEN 'stock_out'        THEN (SELECT user_id FROM stock_outs        WHERE id = m.reference_id)
		               WHEN 'stock_adjustment' THEN (SELECT user_id FROM stock_adjustments WHERE id = m.reference_id)
		               WHEN 'stock_transfer'   THEN (SELECT user_id FROM stock_transfers   WHERE id = m.reference_id)
		           END)
		FROM stock_movements m
		JOIN barang b ON m.barang_id = b.id
		JOIN lokasi l ON m.lokasi_id = l.id
		ORDER BY m.id DESC
		LIMIT ?`, limit)
	if err != nil {
		return []model.AktivitasRow{}
	}
	defer rows.Close()

	list := []model.AktivitasRow{}
	for rows.Next() {
		var x model.AktivitasRow
		if err := rows.Scan(&x.Waktu, &x.Type, &x.Referensi, &x.NamaBarang, &x.NamaLokasi,
			&x.Quantity, &x.StockAfter, &x.NamaUser); err == nil {
			list = append(list, x)
		}
	}
	return list
}
