package model

type StockTransferDetail struct {
	ID              int     `json:"id"`
	StockTransferID int     `json:"stock_transfer_id"`
	BarangID        int     `json:"barang_id"`
	NamaBarang      *string `json:"nama_barang,omitempty"`
	Quantity        int     `json:"quantity"`
}

type StockTransfer struct {
	ID           int                   `json:"id"`
	Code         string                `json:"code"`
	FromLokasiID int                   `json:"from_lokasi_id"`
	NamaFrom     *string               `json:"nama_from,omitempty"`
	ToLokasiID   int                   `json:"to_lokasi_id"`
	NamaTo       *string               `json:"nama_to,omitempty"`
	UserID       *int                  `json:"user_id"`
	NamaUser     *string               `json:"nama_user,omitempty"`
	Tanggal      string                `json:"tanggal"`
	Catatan      string                `json:"catatan"`
	Status       string                `json:"status"`
	TotalQty     int                   `json:"total_qty"`
	Details      []StockTransferDetail `json:"details"`
}
