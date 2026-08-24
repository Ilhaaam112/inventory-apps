package model

type Supplier struct {
	ID           int    `json:"id"`
	NamaSupplier string `json:"nama_supplier"`
	Kontak       string `json:"kontak"`
	Alamat       string `json:"alamat"`
}
