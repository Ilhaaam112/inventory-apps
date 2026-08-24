package model

// KartuStokRow = satu pergerakan pada kartu stok.
type KartuStokRow struct {
	Tanggal      string `json:"tanggal"`
	Type         string `json:"type"`
	Referensi    string `json:"referensi"`
	Masuk        int    `json:"masuk"`
	Keluar       int    `json:"keluar"`
	SaldoSebelum int    `json:"saldo_sebelum"`
	SaldoSesudah int    `json:"saldo_sesudah"`
}

// KartuStok selalu terikat pada satu barang di satu gudang, karena
// saldo berjalan hanya bermakna kalau gudangnya tunggal.
type KartuStok struct {
	BarangID   int            `json:"barang_id"`
	NamaBarang string         `json:"nama_barang"`
	NamaSatuan *string        `json:"nama_satuan,omitempty"`
	LokasiID   int            `json:"lokasi_id"`
	NamaLokasi string         `json:"nama_lokasi"`
	Start      string         `json:"start"`
	End        string         `json:"end"`
	SaldoAwal  int            `json:"saldo_awal"`
	SaldoAkhir int            `json:"saldo_akhir"`
	Rows       []KartuStokRow `json:"rows"`
}
