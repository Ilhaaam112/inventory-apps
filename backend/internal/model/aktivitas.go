package model

// AktivitasRow = satu baris riwayat pergerakan stok terbaru,
// dilengkapi nomor transaksi dan nama petugas yang menyimpannya.
type AktivitasRow struct {
	Waktu      string  `json:"waktu"`
	Type       string  `json:"type"`
	Referensi  string  `json:"referensi"`
	NamaBarang string  `json:"nama_barang"`
	NamaLokasi string  `json:"nama_lokasi"`
	Quantity   int     `json:"quantity"`
	StockAfter int     `json:"stock_after"`
	NamaUser   *string `json:"nama_user,omitempty"`
}
