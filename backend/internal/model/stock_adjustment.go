package model

type StockAdjustmentDetail struct {
	ID                int     `json:"id"`
	StockAdjustmentID int     `json:"stock_adjustment_id"`
	BarangID          int     `json:"barang_id"`
	NamaBarang        *string `json:"nama_barang,omitempty"`
	SystemStock       int     `json:"system_stock"`
	ActualStock       int     `json:"actual_stock"`
	Difference        int     `json:"difference"`
}

type StockAdjustment struct {
	ID         int                     `json:"id"`
	Code       string                  `json:"code"`
	LokasiID   int                     `json:"lokasi_id"`
	NamaLokasi *string                 `json:"nama_lokasi,omitempty"`
	UserID     *int                    `json:"user_id"`
	NamaUser   *string                 `json:"nama_user,omitempty"`
	Tanggal    string                  `json:"tanggal"`
	Alasan     string                  `json:"alasan"`
	Status     string                  `json:"status"`
	TotalItem  int                     `json:"total_item"`
	Details    []StockAdjustmentDetail `json:"details"`
}
