package model

// LaporanMasukRow = rincian barang masuk dalam satu periode.
type LaporanMasukRow struct {
	Tanggal      string  `json:"tanggal"`
	Code         string  `json:"code"`
	NamaSupplier *string `json:"nama_supplier,omitempty"`
	NamaLokasi   string  `json:"nama_lokasi"`
	NamaBarang   string  `json:"nama_barang"`
	NamaSatuan   *string `json:"nama_satuan,omitempty"`
	Quantity     int     `json:"quantity"`
	HargaBeli    float64 `json:"harga_beli"`
	Subtotal     float64 `json:"subtotal"`
}
