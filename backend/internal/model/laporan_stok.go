package model

// LaporanStokRow = posisi stok saat ini, satu baris per barang per gudang.
type LaporanStokRow struct {
	LokasiID     int     `json:"lokasi_id"`
	NamaLokasi   string  `json:"nama_lokasi"`
	BarangID     int     `json:"barang_id"`
	NamaBarang   string  `json:"nama_barang"`
	NamaKategori *string `json:"nama_kategori,omitempty"`
	NamaSatuan   *string `json:"nama_satuan,omitempty"`
	Quantity     int     `json:"quantity"`
	Harga        float64 `json:"harga"`
	Nilai        float64 `json:"nilai"`
}
