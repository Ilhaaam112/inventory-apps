package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type SatuanRepository struct {
	db *sql.DB
}

func NewSatuanRepository(db *sql.DB) *SatuanRepository {
	return &SatuanRepository{db: db}
}

func (r *SatuanRepository) GetAll() []model.Satuan {
	rows, err := r.db.Query("SELECT id, nama_satuan, keterangan FROM satuan ORDER BY id DESC")
	if err != nil {
		return []model.Satuan{}
	}
	defer rows.Close()

	var list []model.Satuan
	for rows.Next() {
		var s model.Satuan
		rows.Scan(&s.ID, &s.NamaSatuan, &s.Keterangan)
		list = append(list, s)
	}
	return list
}

func (r *SatuanRepository) GetByID(id int) (model.Satuan, error) {
	var s model.Satuan
	row := r.db.QueryRow("SELECT id, nama_satuan, keterangan FROM satuan WHERE id = ?", id)
	if err := row.Scan(&s.ID, &s.NamaSatuan, &s.Keterangan); err != nil {
		return model.Satuan{}, errors.New("satuan tidak ditemukan")
	}
	return s, nil
}

func (r *SatuanRepository) Create(s model.Satuan) model.Satuan {
	result, err := r.db.Exec("INSERT INTO satuan (nama_satuan, keterangan) VALUES (?, ?)", s.NamaSatuan, s.Keterangan)
	if err != nil {
		return model.Satuan{}
	}
	id, _ := result.LastInsertId()
	s.ID = int(id)
	return s
}

func (r *SatuanRepository) Update(id int, updated model.Satuan) (model.Satuan, error) {
	_, err := r.db.Exec("UPDATE satuan SET nama_satuan = ?, keterangan = ? WHERE id = ?", updated.NamaSatuan, updated.Keterangan, id)
	if err != nil {
		return model.Satuan{}, err
	}
	updated.ID = id
	return updated, nil
}

func (r *SatuanRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM satuan WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("satuan tidak ditemukan")
	}
	return nil
}
