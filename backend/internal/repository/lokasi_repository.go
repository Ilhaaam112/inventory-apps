package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type LokasiRepository struct {
	db *sql.DB
}

func NewLokasiRepository(db *sql.DB) *LokasiRepository {
	return &LokasiRepository{db: db}
}

func (r *LokasiRepository) GetAll() []model.Lokasi {
	rows, err := r.db.Query("SELECT id, nama_lokasi, keterangan FROM lokasi ORDER BY id DESC")
	if err != nil {
		return []model.Lokasi{}
	}
	defer rows.Close()

	var list []model.Lokasi
	for rows.Next() {
		var l model.Lokasi
		rows.Scan(&l.ID, &l.NamaLokasi, &l.Keterangan)
		list = append(list, l)
	}
	return list
}

func (r *LokasiRepository) GetByID(id int) (model.Lokasi, error) {
	var l model.Lokasi
	row := r.db.QueryRow("SELECT id, nama_lokasi, keterangan FROM lokasi WHERE id = ?", id)
	if err := row.Scan(&l.ID, &l.NamaLokasi, &l.Keterangan); err != nil {
		return model.Lokasi{}, errors.New("lokasi tidak ditemukan")
	}
	return l, nil
}

func (r *LokasiRepository) Create(l model.Lokasi) model.Lokasi {
	result, err := r.db.Exec("INSERT INTO lokasi (nama_lokasi, keterangan) VALUES (?, ?)", l.NamaLokasi, l.Keterangan)
	if err != nil {
		return model.Lokasi{}
	}
	id, _ := result.LastInsertId()
	l.ID = int(id)
	return l
}

func (r *LokasiRepository) Update(id int, updated model.Lokasi) (model.Lokasi, error) {
	_, err := r.db.Exec("UPDATE lokasi SET nama_lokasi = ?, keterangan = ? WHERE id = ?", updated.NamaLokasi, updated.Keterangan, id)
	if err != nil {
		return model.Lokasi{}, err
	}
	updated.ID = id
	return updated, nil
}

func (r *LokasiRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM lokasi WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("lokasi tidak ditemukan")
	}
	return nil
}
