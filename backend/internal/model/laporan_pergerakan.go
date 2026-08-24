package model

// LaporanPergerakanRow = rekap mutasi stok per barang per gudang:
// saldo awal, total masuk, total keluar, penyesuaian, saldo akhir.
type LaporanPergerakanRow struct {
	LokasiID    int     `json:"lokasi_id"`
	NamaLokasi  string  `json:"nama_lokasi"`
	BarangID    int     `json:"barang_id"`
	NamaBarang  string  `json:"nama_barang"`
	NamaSatuan  *string `json:"nama_satuan,omitempty"`
	SaldoAwal   int     `json:"saldo_awal"`
	Masuk       int     `json:"masuk"`
	Keluar      int     `json:"keluar"`
	Penyesuaian int     `json:"penyesuaian"`
	SaldoAkhir  int     `json:"saldo_akhir"`
}
