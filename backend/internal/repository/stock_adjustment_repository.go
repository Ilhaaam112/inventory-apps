package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type StockAdjustmentRepository struct {
	db *sql.DB
}

func NewStockAdjustmentRepository(db *sql.DB) *StockAdjustmentRepository {
	return &StockAdjustmentRepository{db: db}
}

const adjustmentHeaderQuery = `
	SELECT a.id, a.code, a.lokasi_id, l.nama_lokasi, a.user_id, u.nama_lengkap,
	       DATE_FORMAT(a.tanggal, '%Y-%m-%d'), a.alasan, a.status,
	       COALESCE((SELECT COUNT(*) FROM stock_adjustment_details d WHERE d.stock_adjustment_id = a.id), 0)
	FROM stock_adjustments a
	LEFT JOIN lokasi l ON a.lokasi_id = l.id
	LEFT JOIN users  u ON a.user_id   = u.id
`

func scanAdjustment(row rowScanner) (model.StockAdjustment, error) {
	var a model.StockAdjustment
	err := row.Scan(&a.ID, &a.Code, &a.LokasiID, &a.NamaLokasi, &a.UserID, &a.NamaUser,
		&a.Tanggal, &a.Alasan, &a.Status, &a.TotalItem)
	return a, err
}

func (r *StockAdjustmentRepository) GetAll() []model.StockAdjustment {
	rows, err := r.db.Query(adjustmentHeaderQuery + " ORDER BY a.id DESC")
	if err != nil {
		return []model.StockAdjustment{}
	}
	defer rows.Close()

	list := []model.StockAdjustment{}
	for rows.Next() {
		if a, err := scanAdjustment(rows); err == nil {
			list = append(list, a)
		}
	}
	return list
}

func (r *StockAdjustmentRepository) GetByID(id int) (model.StockAdjustment, error) {
	a, err := scanAdjustment(r.db.QueryRow(adjustmentHeaderQuery+" WHERE a.id = ?", id))
	if err != nil {
		return model.StockAdjustment{}, errors.New("transaksi penyesuaian tidak ditemukan")
	}

	rows, err := r.db.Query(`
		SELECT d.id, d.stock_adjustment_id, d.barang_id, b.nama, d.system_stock, d.actual_stock, d.difference
		FROM stock_adjustment_details d
		LEFT JOIN barang b ON d.barang_id = b.id
		WHERE d.stock_adjustment_id = ? ORDER BY d.id`, id)
	if err != nil {
		return a, nil
	}
	defer rows.Close()

	a.Details = []model.StockAdjustmentDetail{}
	for rows.Next() {
		var d model.StockAdjustmentDetail
		if err := rows.Scan(&d.ID, &d.StockAdjustmentID, &d.BarangID, &d.NamaBarang,
			&d.SystemStock, &d.ActualStock, &d.Difference); err == nil {
			a.Details = append(a.Details, d)
		}
	}
	return a, nil
}

func (r *StockAdjustmentRepository) Create(a model.StockAdjustment) (model.StockAdjustment, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.StockAdjustment{}, err
	}
	defer tx.Rollback()

	if err := ensureExists(tx, "lokasi", a.LokasiID, "gudang tidak ditemukan"); err != nil {
		return model.StockAdjustment{}, err
	}

	code, err := generateCode(tx, "stock_adjustments", "ADJ")
	if err != nil {
		return model.StockAdjustment{}, err
	}

	res, err := tx.Exec(`
		INSERT INTO stock_adjustments (code, lokasi_id, user_id, tanggal, alasan, status)
		VALUES (?, ?, ?, ?, ?, 'POSTED')`,
		code, a.LokasiID, a.UserID, a.Tanggal, a.Alasan)
	if err != nil {
		return model.StockAdjustment{}, err
	}
	newID64, _ := res.LastInsertId()
	newID := int(newID64)

	for _, d := range a.Details {
		if err := ensureExists(tx, "barang", d.BarangID, "barang tidak ditemukan"); err != nil {
			return model.StockAdjustment{}, err
		}

		// Stok sistem selalu diambil ulang dari database, bukan dari frontend.
		systemStock, err := currentStock(tx, a.LokasiID, d.BarangID)
		if err != nil {
			return model.StockAdjustment{}, err
		}
		difference := d.ActualStock - systemStock

		if _, err := tx.Exec(`
			INSERT INTO stock_adjustment_details
				(stock_adjustment_id, barang_id, system_stock, actual_stock, difference)
			VALUES (?, ?, ?, ?, ?)`,
			newID, d.BarangID, systemStock, d.ActualStock, difference,
		); err != nil {
			return model.StockAdjustment{}, err
		}

		if difference != 0 {
			if err := applyMovement(tx, a.LokasiID, d.BarangID, difference, "ADJUSTMENT", "stock_adjustment", newID); err != nil {
				return model.StockAdjustment{}, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return model.StockAdjustment{}, err
	}
	return r.GetByID(newID)
}
