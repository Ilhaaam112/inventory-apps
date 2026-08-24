package repository

import (
	"database/sql"

	"github.com/username/belajar_go/backend/internal/model"
)

type WarehouseStockRepository struct {
	db *sql.DB
}

func NewWarehouseStockRepository(db *sql.DB) *WarehouseStockRepository {
	return &WarehouseStockRepository{db: db}
}

// GetAll menerima filter opsional: kirim 0 untuk mengabaikan filter tersebut.
func (r *WarehouseStockRepository) GetAll(lokasiID, barangID int) []model.WarehouseStock {
	rows, err := r.db.Query(`
		SELECT ws.lokasi_id, l.nama_lokasi, ws.barang_id, b.nama, s.nama_satuan, ws.quantity
		FROM warehouse_stocks ws
		JOIN barang b      ON ws.barang_id = b.id
		JOIN lokasi l      ON ws.lokasi_id = l.id
		LEFT JOIN satuan s ON b.satuan_id  = s.id
		WHERE (? = 0 OR ws.lokasi_id = ?) AND (? = 0 OR ws.barang_id = ?)
		ORDER BY l.nama_lokasi, b.nama`,
		lokasiID, lokasiID, barangID, barangID)
	if err != nil {
		return []model.WarehouseStock{}
	}
	defer rows.Close()

	list := []model.WarehouseStock{}
	for rows.Next() {
		var w model.WarehouseStock
		if err := rows.Scan(&w.LokasiID, &w.NamaLokasi, &w.BarangID, &w.NamaBarang, &w.NamaSatuan, &w.Quantity); err == nil {
			list = append(list, w)
		}
	}
	return list
}
