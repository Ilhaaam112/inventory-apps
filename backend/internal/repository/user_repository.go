package repository

import (
	"database/sql"
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (model.User, error) {
	var u model.User
	row := r.db.QueryRow("SELECT id, username, password, nama_lengkap FROM users WHERE username = ?", username)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.NamaLengkap)
	if err != nil {
		return model.User{}, errors.New("username tidak ditemukan")
	}
	return u, nil
}




// ---------- BARU: untuk halaman Profil ----------

func (r *UserRepository) GetByID(id int) (model.User, error) {
	var u model.User
	row := r.db.QueryRow("SELECT id, username, password, nama_lengkap FROM users WHERE id = ?", id)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.NamaLengkap)
	if err != nil {
		return model.User{}, errors.New("user tidak ditemukan")
	}
	return u, nil
}

func (r *UserRepository) UpdateProfile(id int, namaLengkap string) error {
	_, err := r.db.Exec("UPDATE users SET nama_lengkap = ? WHERE id = ?", namaLengkap, id)
	return err
}

func (r *UserRepository) UpdatePassword(id int, hashedPassword string) error {
	_, err := r.db.Exec("UPDATE users SET password = ? WHERE id = ?", hashedPassword, id)
	return err
}