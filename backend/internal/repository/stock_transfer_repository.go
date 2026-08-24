package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type StockTransferRepository struct {
	db *sql.DB
}

func NewStockTransferRepository(db *sql.DB) *StockTransferRepository {
	return &StockTransferRepository{db: db}
}

const transferHeaderQuery = `
	SELECT t.id, t.code, t.from_lokasi_id, lf.nama_lokasi, t.to_lokasi_id, lt.nama_lokasi,
	       t.user_id, u.nama_lengkap, DATE_FORMAT(t.tanggal, '%Y-%m-%d'), t.catatan, t.status,
	       COALESCE((SELECT SUM(d.quantity) FROM stock_transfer_details d WHERE d.stock_transfer_id = t.id), 0)
	FROM stock_transfers t
	LEFT JOIN lokasi lf ON t.from_lokasi_id = lf.id
	LEFT JOIN lokasi lt ON t.to_lokasi_id   = lt.id
	LEFT JOIN users  u  ON t.user_id        = u.id
`

func scanTransfer(row rowScanner) (model.StockTransfer, error) {
	var t model.StockTransfer
	err := row.Scan(&t.ID, &t.Code, &t.FromLokasiID, &t.NamaFrom, &t.ToLokasiID, &t.NamaTo,
		&t.UserID, &t.NamaUser, &t.Tanggal, &t.Catatan, &t.Status, &t.TotalQty)
	return t, err
}

func (r *StockTransferRepository) GetAll() []model.StockTransfer {
	rows, err := r.db.Query(transferHeaderQuery + " ORDER BY t.id DESC")
	if err != nil {
		return []model.StockTransfer{}
	}
	defer rows.Close()

	list := []model.StockTransfer{}
	for rows.Next() {
		if t, err := scanTransfer(rows); err == nil {
			list = append(list, t)
		}
	}
	return list
}

func (r *StockTransferRepository) GetByID(id int) (model.StockTransfer, error) {
	t, err := scanTransfer(r.db.QueryRow(transferHeaderQuery+" WHERE t.id = ?", id))
	if err != nil {
		return model.StockTransfer{}, errors.New("transaksi transfer tidak ditemukan")
	}

	rows, err := r.db.Query(`
		SELECT d.id, d.stock_transfer_id, d.barang_id, b.nama, d.quantity
		FROM stock_transfer_details d
		LEFT JOIN barang b ON d.barang_id = b.id
		WHERE d.stock_transfer_id = ? ORDER BY d.id`, id)
	if err != nil {
		return t, nil
	}
	defer rows.Close()

	t.Details = []model.StockTransferDetail{}
	for rows.Next() {
		var d model.StockTransferDetail
		if err := rows.Scan(&d.ID, &d.StockTransferID, &d.BarangID, &d.NamaBarang, &d.Quantity); err == nil {
			t.Details = append(t.Details, d)
		}
	}
	return t, nil
}

func (r *StockTransferRepository) Create(t model.StockTransfer) (model.StockTransfer, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.StockTransfer{}, err
	}
	defer tx.Rollback()

	if err := ensureExists(tx, "lokasi", t.FromLokasiID, "gudang asal tidak ditemukan"); err != nil {
		return model.StockTransfer{}, err
	}
	if err := ensureExists(tx, "lokasi", t.ToLokasiID, "gudang tujuan tidak ditemukan"); err != nil {
		return model.StockTransfer{}, err
	}

	code, err := generateCode(tx, "stock_transfers", "TRF")
	if err != nil {
		return model.StockTransfer{}, err
	}

	res, err := tx.Exec(`
		INSERT INTO stock_transfers (code, from_lokasi_id, to_lokasi_id, user_id, tanggal, catatan, status)
		VALUES (?, ?, ?, ?, ?, ?, 'POSTED')`,
		code, t.FromLokasiID, t.ToLokasiID, t.UserID, t.Tanggal, t.Catatan)
	if err != nil {
		return model.StockTransfer{}, err
	}
	newID64, _ := res.LastInsertId()
	newID := int(newID64)

	for _, d := range t.Details {
		if err := ensureExists(tx, "barang", d.BarangID, "barang tidak ditemukan"); err != nil {
			return model.StockTransfer{}, err
		}
		if _, err := tx.Exec(
			"INSERT INTO stock_transfer_details (stock_transfer_id, barang_id, quantity) VALUES (?, ?, ?)",
			newID, d.BarangID, d.Quantity,
		); err != nil {
			return model.StockTransfer{}, err
		}
		// Satu transfer menghasilkan dua pergerakan.
		if err := applyMovement(tx, t.FromLokasiID, d.BarangID, -d.Quantity, "TRANSFER_OUT", "stock_transfer", newID); err != nil {
			return model.StockTransfer{}, err
		}
		if err := applyMovement(tx, t.ToLokasiID, d.BarangID, d.Quantity, "TRANSFER_IN", "stock_transfer", newID); err != nil {
			return model.StockTransfer{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return model.StockTransfer{}, err
	}
	return r.GetByID(newID)
}
