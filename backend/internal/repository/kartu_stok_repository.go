package repository

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

// KartuStok menampilkan riwayat pergerakan satu barang di satu gudang,
// lengkap dengan saldo sebelum dan sesudah tiap transaksi.
func (r *LaporanRepository) KartuStok(barangID, lokasiID int, start, end string) (model.KartuStok, error) {
	var k model.KartuStok
	k.BarangID = barangID
	k.LokasiID = lokasiID
	k.Start = start
	k.End = end

	// Header: nama barang, satuan, gudang, dan stok terkini.
	var stokSekarang int
	err := r.db.QueryRow(`
		SELECT b.nama, s.nama_satuan, l.nama_lokasi, COALESCE(ws.quantity, 0)
		FROM barang b
		JOIN lokasi l                 ON l.id = ?
		LEFT JOIN satuan s            ON b.satuan_id = s.id
		LEFT JOIN warehouse_stocks ws ON ws.barang_id = b.id AND ws.lokasi_id = l.id
		WHERE b.id = ?`, lokasiID, barangID,
	).Scan(&k.NamaBarang, &k.NamaSatuan, &k.NamaLokasi, &stokSekarang)
	if err != nil {
		return model.KartuStok{}, errors.New("barang atau gudang tidak ditemukan")
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
		       m.quantity, m.stock_before, m.stock_after
		FROM stock_movements m
		WHERE m.barang_id = ? AND m.lokasi_id = ?
		  AND DATE(m.created_at) BETWEEN ? AND ?
		ORDER BY m.id ASC`, barangID, lokasiID, start, end)
	if err != nil {
		return k, nil
	}
	defer rows.Close()

	k.Rows = []model.KartuStokRow{}
	for rows.Next() {
		var x model.KartuStokRow
		var qty int
		if err := rows.Scan(&x.Tanggal, &x.Type, &x.Referensi, &qty, &x.SaldoSebelum, &x.SaldoSesudah); err != nil {
			continue
		}
		if qty >= 0 {
			x.Masuk = qty
		} else {
			x.Keluar = -qty
		}
		k.Rows = append(k.Rows, x)
	}

	if len(k.Rows) > 0 {
		k.SaldoAwal = k.Rows[0].SaldoSebelum
		k.SaldoAkhir = k.Rows[len(k.Rows)-1].SaldoSesudah
		return k, nil
	}

	// Tidak ada pergerakan dalam periode: saldo awal = saldo akhir,
	// dihitung mundur dari stok sekarang dikurangi pergerakan setelah periode.
	var setelah int
	r.db.QueryRow(`
		SELECT COALESCE(SUM(quantity), 0) FROM stock_movements
		WHERE barang_id = ? AND lokasi_id = ? AND DATE(created_at) > ?`,
		barangID, lokasiID, end).Scan(&setelah)

	k.SaldoAwal = stokSekarang - setelah
	k.SaldoAkhir = k.SaldoAwal
	return k, nil
}
