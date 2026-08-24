package model

type StockInDetail struct {
	ID         int     `json:"id"`
	StockInID  int     `json:"stock_in_id"`
	BarangID   int     `json:"barang_id"`
	NamaBarang *string `json:"nama_barang,omitempty"`
	Quantity   int     `json:"quantity"`
	HargaBeli  float64 `json:"harga_beli"`
}

type StockIn struct {
	ID           int             `json:"id"`
	Code         string          `json:"code"`
	SupplierID   *int            `json:"supplier_id"`
	NamaSupplier *string         `json:"nama_supplier,omitempty"`
	LokasiID     int             `json:"lokasi_id"`
	NamaLokasi   *string         `json:"nama_lokasi,omitempty"`
	UserID       *int            `json:"user_id"`
	NamaUser     *string         `json:"nama_user,omitempty"`
	Tanggal      string          `json:"tanggal"`
	Catatan      string          `json:"catatan"`
	Status       string          `json:"status"`
	TotalQty     int             `json:"total_qty"`
	Details      []StockInDetail `json:"details"`
}
