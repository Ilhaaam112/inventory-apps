package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type StockInRepository struct {
	db *sql.DB
}

func NewStockInRepository(db *sql.DB) *StockInRepository {
	return &StockInRepository{db: db}
}

const stockInHeaderQuery = `
	SELECT si.id, si.code, si.supplier_id, sp.nama_supplier, si.lokasi_id, l.nama_lokasi,
	       si.user_id, u.nama_lengkap, DATE_FORMAT(si.tanggal, '%Y-%m-%d'),
	       si.catatan, si.status,
	       COALESCE((SELECT SUM(d.quantity) FROM stock_in_details d WHERE d.stock_in_id = si.id), 0)
	FROM stock_ins si
	LEFT JOIN supplier sp ON si.supplier_id = sp.id
	LEFT JOIN lokasi   l  ON si.lokasi_id   = l.id
	LEFT JOIN users    u  ON si.user_id     = u.id
`

func scanStockIn(row rowScanner) (model.StockIn, error) {
	var s model.StockIn
	err := row.Scan(&s.ID, &s.Code, &s.SupplierID, &s.NamaSupplier, &s.LokasiID, &s.NamaLokasi,
		&s.UserID, &s.NamaUser, &s.Tanggal, &s.Catatan, &s.Status, &s.TotalQty)
	return s, err
}

func (r *StockInRepository) GetAll() []model.StockIn {
	rows, err := r.db.Query(stockInHeaderQuery + " ORDER BY si.id DESC")
	if err != nil {
		return []model.StockIn{}
	}
	defer rows.Close()

	list := []model.StockIn{}
	for rows.Next() {
		if s, err := scanStockIn(rows); err == nil {
			list = append(list, s)
		}
	}
	return list
}

func (r *StockInRepository) GetByID(id int) (model.StockIn, error) {
	s, err := scanStockIn(r.db.QueryRow(stockInHeaderQuery+" WHERE si.id = ?", id))
	if err != nil {
		return model.StockIn{}, errors.New("transaksi barang masuk tidak ditemukan")
	}

	rows, err := r.db.Query(`
		SELECT d.id, d.stock_in_id, d.barang_id, b.nama, d.quantity, d.harga_beli
		FROM stock_in_details d
		LEFT JOIN barang b ON d.barang_id = b.id
		WHERE d.stock_in_id = ? ORDER BY d.id`, id)
	if err != nil {
		return s, nil
	}
	defer rows.Close()

	s.Details = []model.StockInDetail{}
	for rows.Next() {
		var d model.StockInDetail
		if err := rows.Scan(&d.ID, &d.StockInID, &d.BarangID, &d.NamaBarang, &d.Quantity, &d.HargaBeli); err == nil {
			s.Details = append(s.Details, d)
		}
	}
	return s, nil
}

func (r *StockInRepository) Create(in model.StockIn) (model.StockIn, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.StockIn{}, err
	}
	defer tx.Rollback()

	if err := ensureExists(tx, "lokasi", in.LokasiID, "gudang tidak ditemukan"); err != nil {
		return model.StockIn{}, err
	}
	if in.SupplierID != nil {
		if err := ensureExists(tx, "supplier", *in.SupplierID, "supplier tidak ditemukan"); err != nil {
			return model.StockIn{}, err
		}
	}

	code, err := generateCode(tx, "stock_ins", "IN")
	if err != nil {
		return model.StockIn{}, err
	}

	res, err := tx.Exec(`
		INSERT INTO stock_ins (code, supplier_id, lokasi_id, user_id, tanggal, catatan, status)
		VALUES (?, ?, ?, ?, ?, ?, 'POSTED')`,
		code, in.SupplierID, in.LokasiID, in.UserID, in.Tanggal, in.Catatan)
	if err != nil {
		return model.StockIn{}, err
	}
	newID64, _ := res.LastInsertId()
	newID := int(newID64)

	for _, d := range in.Details {
		if err := ensureExists(tx, "barang", d.BarangID, "barang tidak ditemukan"); err != nil {
			return model.StockIn{}, err
		}
		if _, err := tx.Exec(
			"INSERT INTO stock_in_details (stock_in_id, barang_id, quantity, harga_beli) VALUES (?, ?, ?, ?)",
			newID, d.BarangID, d.Quantity, d.HargaBeli,
		); err != nil {
			return model.StockIn{}, err
		}
		if err := applyMovement(tx, in.LokasiID, d.BarangID, d.Quantity, "IN", "stock_in", newID); err != nil {
			return model.StockIn{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return model.StockIn{}, err
	}
	return r.GetByID(newID)
}
