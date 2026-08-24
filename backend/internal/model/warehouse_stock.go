package model

// WarehouseStock adalah stok satu barang di satu gudang.
type WarehouseStock struct {
	LokasiID   int     `json:"lokasi_id"`
	NamaLokasi string  `json:"nama_lokasi"`
	BarangID   int     `json:"barang_id"`
	NamaBarang string  `json:"nama_barang"`
	NamaSatuan *string `json:"nama_satuan,omitempty"`
	Quantity   int     `json:"quantity"`
}
