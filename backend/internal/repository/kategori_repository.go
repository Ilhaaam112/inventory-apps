package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type KategoriRepository struct {
	db *sql.DB
}

func NewKategoriRepository(db *sql.DB) *KategoriRepository {
	return &KategoriRepository{db: db}
}

func (r *KategoriRepository) GetAll() []model.Kategori {
	rows, err := r.db.Query("SELECT id, nama_kategori FROM kategori ORDER BY nama_kategori")
	if err != nil {
		return []model.Kategori{}
	}
	defer rows.Close()

	var list []model.Kategori
	for rows.Next() {
		var k model.Kategori
		rows.Scan(&k.ID, &k.NamaKategori)
		list = append(list, k)
	}
	return list
}

func (r *KategoriRepository) GetByID(id int) (model.Kategori, error) {
	var k model.Kategori
	row := r.db.QueryRow("SELECT id, nama_kategori FROM kategori WHERE id = ?", id)
	if err := row.Scan(&k.ID, &k.NamaKategori); err != nil {
		return model.Kategori{}, errors.New("kategori tidak ditemukan")
	}
	return k, nil
}

func (r *KategoriRepository) Create(k model.Kategori) model.Kategori {
	result, err := r.db.Exec("INSERT INTO kategori (nama_kategori) VALUES (?)", k.NamaKategori)
	if err != nil {
		return model.Kategori{}
	}
	id, _ := result.LastInsertId()
	k.ID = int(id)
	return k
}

func (r *KategoriRepository) Update(id int, updated model.Kategori) (model.Kategori, error) {
	_, err := r.db.Exec("UPDATE kategori SET nama_kategori = ? WHERE id = ?", updated.NamaKategori, id)
	if err != nil {
		return model.Kategori{}, err
	}
	updated.ID = id
	return updated, nil
}

func (r *KategoriRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM kategori WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("kategori tidak ditemukan")
	}
	return nil
}