package model

// StokMenipisRow = barang yang stoknya sudah menyentuh atau melewati
// ambang batas minimum di sebuah gudang.
type StokMenipisRow struct {
	LokasiID    int     `json:"lokasi_id"`
	NamaLokasi  string  `json:"nama_lokasi"`
	BarangID    int     `json:"barang_id"`
	NamaBarang  string  `json:"nama_barang"`
	NamaSatuan  *string `json:"nama_satuan,omitempty"`
	Quantity    int     `json:"quantity"`
	StokMinimum int     `json:"stok_minimum"`
	Kekurangan  int     `json:"kekurangan"`
	Habis       bool    `json:"habis"`
}
