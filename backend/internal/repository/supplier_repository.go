package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func (r *SupplierRepository) GetAll() []model.Supplier {
	rows, err := r.db.Query("SELECT id, nama_supplier, kontak, alamat FROM supplier ORDER BY id DESC")
	if err != nil {
		return []model.Supplier{}
	}
	defer rows.Close()

	var list []model.Supplier
	for rows.Next() {
		var s model.Supplier
		rows.Scan(&s.ID, &s.NamaSupplier, &s.Kontak, &s.Alamat)
		list = append(list, s)
	}
	return list
}

func (r *SupplierRepository) GetByID(id int) (model.Supplier, error) {
	var s model.Supplier
	row := r.db.QueryRow("SELECT id, nama_supplier, kontak, alamat FROM supplier WHERE id = ?", id)
	if err := row.Scan(&s.ID, &s.NamaSupplier, &s.Kontak, &s.Alamat); err != nil {
		return model.Supplier{}, errors.New("supplier tidak ditemukan")
	}
	return s, nil
}

func (r *SupplierRepository) Create(s model.Supplier) model.Supplier {
	result, err := r.db.Exec(
		"INSERT INTO supplier (nama_supplier, kontak, alamat) VALUES (?, ?, ?)",
		s.NamaSupplier, s.Kontak, s.Alamat,
	)
	if err != nil {
		return model.Supplier{}
	}
	id, _ := result.LastInsertId()
	s.ID = int(id)
	return s
}

func (r *SupplierRepository) Update(id int, updated model.Supplier) (model.Supplier, error) {
	_, err := r.db.Exec(
		"UPDATE supplier SET nama_supplier = ?, kontak = ?, alamat = ? WHERE id = ?",
		updated.NamaSupplier, updated.Kontak, updated.Alamat, id,
	)
	if err != nil {
		return model.Supplier{}, err
	}
	updated.ID = id
	return updated, nil
}

func (r *SupplierRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM supplier WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("supplier tidak ditemukan")
	}
	return nil
}
