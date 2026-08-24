package repository

import (
	"database/sql"

	"github.com/username/belajar_go/backend/internal/model"
)

type StockMovementRepository struct {
	db *sql.DB
}

func NewStockMovementRepository(db *sql.DB) *StockMovementRepository {
	return &StockMovementRepository{db: db}
}

// GetAll menerima filter opsional: kirim 0 untuk mengabaikan filter tersebut.
func (r *StockMovementRepository) GetAll(lokasiID, barangID, limit int) []model.StockMovement {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := r.db.Query(`
		SELECT m.id, DATE_FORMAT(m.created_at, '%Y-%m-%d %H:%i'), m.barang_id, b.nama,
		       m.lokasi_id, l.nama_lokasi, m.type, m.reference_type, m.reference_id,
		       m.quantity, m.stock_before, m.stock_after
		FROM stock_movements m
		JOIN barang b ON m.barang_id = b.id
		JOIN lokasi l ON m.lokasi_id = l.id
		WHERE (? = 0 OR m.lokasi_id = ?) AND (? = 0 OR m.barang_id = ?)
		ORDER BY m.id DESC LIMIT ?`,
		lokasiID, lokasiID, barangID, barangID, limit)
	if err != nil {
		return []model.StockMovement{}
	}
	defer rows.Close()

	list := []model.StockMovement{}
	for rows.Next() {
		var m model.StockMovement
		if err := rows.Scan(&m.ID, &m.CreatedAt, &m.BarangID, &m.NamaBarang, &m.LokasiID, &m.NamaLokasi,
			&m.Type, &m.ReferenceType, &m.ReferenceID, &m.Quantity, &m.StockBefore, &m.StockAfter); err == nil {
			list = append(list, m)
		}
	}
	return list
}
