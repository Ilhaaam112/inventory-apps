package model

// LaporanKeluarRow = rincian barang keluar dalam satu periode.
type LaporanKeluarRow struct {
	Tanggal    string  `json:"tanggal"`
	Code       string  `json:"code"`
	NamaLokasi string  `json:"nama_lokasi"`
	Tujuan     string  `json:"tujuan"`
	NamaBarang string  `json:"nama_barang"`
	NamaSatuan *string `json:"nama_satuan,omitempty"`
	Quantity   int     `json:"quantity"`
}
