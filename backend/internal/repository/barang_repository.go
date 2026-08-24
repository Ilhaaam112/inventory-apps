package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type BarangRepository struct {
	db *sql.DB
}

func NewBarangRepository(db *sql.DB) *BarangRepository {
	return &BarangRepository{db: db}
}

const baseQuery = `
    SELECT b.id, b.nama, b.harga, b.stok, b.stok_minimum,
           b.kategori_id, k.nama_kategori, b.satuan_id, s.nama_satuan
    FROM barang b
    LEFT JOIN kategori k ON b.kategori_id = k.id
    LEFT JOIN satuan s ON b.satuan_id = s.id
`

func (r *BarangRepository) GetAll() []model.Barang {
	rows, err := r.db.Query(baseQuery + " ORDER BY b.id DESC")
	if err != nil {
		return []model.Barang{}
	}
	defer rows.Close()

	var list []model.Barang
	for rows.Next() {
		var b model.Barang
		rows.Scan(&b.ID, &b.Nama, &b.Harga, &b.Stok, &b.StokMinimum,
			&b.KategoriID, &b.NamaKategori, &b.SatuanID, &b.NamaSatuan)
		list = append(list, b)
	}
	return list
}

func (r *BarangRepository) GetByID(id int) (model.Barang, error) {
	var b model.Barang
	row := r.db.QueryRow(baseQuery+" WHERE b.id = ?", id)
	err := row.Scan(&b.ID, &b.Nama, &b.Harga, &b.Stok, &b.StokMinimum,
		&b.KategoriID, &b.NamaKategori, &b.SatuanID, &b.NamaSatuan)
	if err != nil {
		return model.Barang{}, errors.New("barang tidak ditemukan")
	}
	return b, nil
}

// Create tidak lagi menerima stok awal: stok hanya boleh bertambah
// lewat transaksi Barang Masuk, supaya warehouse_stocks dan
// stock_movements tetap sinkron.
func (r *BarangRepository) Create(b model.Barang) (model.Barang, error) {
	result, err := r.db.Exec(
		"INSERT INTO barang (nama, harga, stok, stok_minimum, kategori_id, satuan_id) VALUES (?, ?, 0, ?, ?, ?)",
		b.Nama, b.Harga, b.StokMinimum, b.KategoriID, b.SatuanID,
	)
	if err != nil {
		return model.Barang{}, bungkusError(err)
	}
	id, _ := result.LastInsertId()
	b.ID = int(id)
	b.Stok = 0
	return b, nil
}

// Update sengaja tidak menyentuh kolom stok.
func (r *BarangRepository) Update(id int, updated model.Barang) (model.Barang, error) {
	_, err := r.db.Exec(
		"UPDATE barang SET nama = ?, harga = ?, stok_minimum = ?, kategori_id = ?, satuan_id = ? WHERE id = ?",
		updated.Nama, updated.Harga, updated.StokMinimum, updated.KategoriID, updated.SatuanID, id,
	)
	if err != nil {
		return model.Barang{}, bungkusError(err)
	}
	return r.GetByID(id)
}

func (r *BarangRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM barang WHERE id = ?", id)
	if err != nil {
		return bungkusError(err)
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("barang tidak ditemukan")
	}
	return nil
}
