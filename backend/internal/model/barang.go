package model

type Barang struct {
	ID           int     `json:"id"`
	Nama         string  `json:"nama"`
	Harga        float64 `json:"harga"`
	Stok         int     `json:"stok"`
	KategoriID   *int    `json:"kategori_id"`
	NamaKategori *string `json:"nama_kategori,omitempty"`
	SatuanID     *int    `json:"satuan_id"`
	NamaSatuan   *string `json:"nama_satuan,omitempty"`
}