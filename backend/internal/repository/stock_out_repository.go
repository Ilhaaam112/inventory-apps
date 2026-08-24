package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type StockOutRepository struct {
	db *sql.DB
}

func NewStockOutRepository(db *sql.DB) *StockOutRepository {
	return &StockOutRepository{db: db}
}

const stockOutHeaderQuery = `
	SELECT so.id, so.code, so.lokasi_id, l.nama_lokasi, so.user_id, u.nama_lengkap,
	       DATE_FORMAT(so.tanggal, '%Y-%m-%d'), so.tujuan, so.catatan, so.status,
	       COALESCE((SELECT SUM(d.quantity) FROM stock_out_details d WHERE d.stock_out_id = so.id), 0)
	FROM stock_outs so
	LEFT JOIN lokasi l ON so.lokasi_id = l.id
	LEFT JOIN users  u ON so.user_id   = u.id
`

func scanStockOut(row rowScanner) (model.StockOut, error) {
	var s model.StockOut
	err := row.Scan(&s.ID, &s.Code, &s.LokasiID, &s.NamaLokasi, &s.UserID, &s.NamaUser,
		&s.Tanggal, &s.Tujuan, &s.Catatan, &s.Status, &s.TotalQty)
	return s, err
}

func (r *StockOutRepository) GetAll() []model.StockOut {
	rows, err := r.db.Query(stockOutHeaderQuery + " ORDER BY so.id DESC")
	if err != nil {
		return []model.StockOut{}
	}
	defer rows.Close()

	list := []model.StockOut{}
	for rows.Next() {
		if s, err := scanStockOut(rows); err == nil {
			list = append(list, s)
		}
	}
	return list
}

func (r *StockOutRepository) GetByID(id int) (model.StockOut, error) {
	s, err := scanStockOut(r.db.QueryRow(stockOutHeaderQuery+" WHERE so.id = ?", id))
	if err != nil {
		return model.StockOut{}, errors.New("transaksi barang keluar tidak ditemukan")
	}

	rows, err := r.db.Query(`
		SELECT d.id, d.stock_out_id, d.barang_id, b.nama, d.quantity
		FROM stock_out_details d
		LEFT JOIN barang b ON d.barang_id = b.id
		WHERE d.stock_out_id = ? ORDER BY d.id`, id)
	if err != nil {
		return s, nil
	}
	defer rows.Close()

	s.Details = []model.StockOutDetail{}
	for rows.Next() {
		var d model.StockOutDetail
		if err := rows.Scan(&d.ID, &d.StockOutID, &d.BarangID, &d.NamaBarang, &d.Quantity); err == nil {
			s.Details = append(s.Details, d)
		}
	}
	return s, nil
}

func (r *StockOutRepository) Create(out model.StockOut) (model.StockOut, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.StockOut{}, err
	}
	defer tx.Rollback()

	if err := ensureExists(tx, "lokasi", out.LokasiID, "gudang tidak ditemukan"); err != nil {
		return model.StockOut{}, err
	}

	code, err := generateCode(tx, "stock_outs", "OUT")
	if err != nil {
		return model.StockOut{}, err
	}

	res, err := tx.Exec(`
		INSERT INTO stock_outs (code, lokasi_id, user_id, tanggal, tujuan, catatan, status)
		VALUES (?, ?, ?, ?, ?, ?, 'POSTED')`,
		code, out.LokasiID, out.UserID, out.Tanggal, out.Tujuan, out.Catatan)
	if err != nil {
		return model.StockOut{}, err
	}
	newID64, _ := res.LastInsertId()
	newID := int(newID64)

	for _, d := range out.Details {
		if err := ensureExists(tx, "barang", d.BarangID, "barang tidak ditemukan"); err != nil {
			return model.StockOut{}, err
		}
		if _, err := tx.Exec(
			"INSERT INTO stock_out_details (stock_out_id, barang_id, quantity) VALUES (?, ?, ?)",
			newID, d.BarangID, d.Quantity,
		); err != nil {
			return model.StockOut{}, err
		}
		// delta negatif -> applyMovement otomatis menolak kalau stok kurang
		if err := applyMovement(tx, out.LokasiID, d.BarangID, -d.Quantity, "OUT", "stock_out", newID); err != nil {
			return model.StockOut{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return model.StockOut{}, err
	}
	return r.GetByID(newID)
}
