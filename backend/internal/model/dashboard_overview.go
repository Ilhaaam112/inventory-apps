package model

// TrenHarian = ringkasan masuk/keluar untuk satu hari.
type TrenHarian struct {
	Tanggal string `json:"tanggal"`
	Masuk   int    `json:"masuk"`
	Keluar  int    `json:"keluar"`
}

// DashboardOverview = angka-angka ringkas untuk halaman utama.
// Tren mengikuti periode yang dipilih user; sisanya selalu kondisi terkini.
type DashboardOverview struct {
	TotalBarang     int          `json:"total_barang"`
	TotalGudang     int          `json:"total_gudang"`
	TotalSupplier   int          `json:"total_supplier"`
	TotalUnit       int          `json:"total_unit"`
	NilaiPersediaan float64      `json:"nilai_persediaan"`
	StokMenipis     int          `json:"stok_menipis"`
	StokHabis       int          `json:"stok_habis"`
	MasukHariIni    int          `json:"masuk_hari_ini"`
	KeluarHariIni   int          `json:"keluar_hari_ini"`
	TransaksiBulan  int          `json:"transaksi_bulan_ini"`
	Start           string       `json:"start"`
	End             string       `json:"end"`
	MasukPeriode    int          `json:"masuk_periode"`
	KeluarPeriode   int          `json:"keluar_periode"`
	Tren            []TrenHarian `json:"tren"`
}
