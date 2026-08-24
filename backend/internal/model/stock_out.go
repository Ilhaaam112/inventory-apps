package model

type StockOutDetail struct {
	ID         int     `json:"id"`
	StockOutID int     `json:"stock_out_id"`
	BarangID   int     `json:"barang_id"`
	NamaBarang *string `json:"nama_barang,omitempty"`
	Quantity   int     `json:"quantity"`
}

type StockOut struct {
	ID         int              `json:"id"`
	Code       string           `json:"code"`
	LokasiID   int              `json:"lokasi_id"`
	NamaLokasi *string          `json:"nama_lokasi,omitempty"`
	UserID     *int             `json:"user_id"`
	NamaUser   *string          `json:"nama_user,omitempty"`
	Tanggal    string           `json:"tanggal"`
	Tujuan     string           `json:"tujuan"`
	Catatan    string           `json:"catatan"`
	Status     string           `json:"status"`
	TotalQty   int              `json:"total_qty"`
	Details    []StockOutDetail `json:"details"`
}
