package model

// StockMovement adalah satu baris kartu stok.
// Quantity disimpan bertanda: positif menambah, negatif mengurangi.
type StockMovement struct {
	ID            int    `json:"id"`
	CreatedAt     string `json:"created_at"`
	BarangID      int    `json:"barang_id"`
	NamaBarang    string `json:"nama_barang"`
	LokasiID      int    `json:"lokasi_id"`
	NamaLokasi    string `json:"nama_lokasi"`
	Type          string `json:"type"`
	ReferenceType string `json:"reference_type"`
	ReferenceID   int    `json:"reference_id"`
	Quantity      int    `json:"quantity"`
	StockBefore   int    `json:"stock_before"`
	StockAfter    int    `json:"stock_after"`
}
