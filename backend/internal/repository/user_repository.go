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

const userBaseQuery = `
	SELECT u.id, u.username, u.password, u.nama_lengkap, u.role_id, COALESCE(r.name, '')
	FROM users u
	LEFT JOIN roles r ON u.role_id = r.id
`

func (r *UserRepository) GetByUsername(username string) (model.User, error) {
	var u model.User
	err := r.db.QueryRow(userBaseQuery+" WHERE u.username = ?", username).
		Scan(&u.ID, &u.Username, &u.Password, &u.NamaLengkap, &u.RoleID, &u.RoleName)
	if err != nil {
		return model.User{}, errors.New("username tidak ditemukan")
	}
	return u, nil
}

func (r *UserRepository) GetByID(id int) (model.User, error) {
	var u model.User
	err := r.db.QueryRow(userBaseQuery+" WHERE u.id = ?", id).
		Scan(&u.ID, &u.Username, &u.Password, &u.NamaLengkap, &u.RoleID, &u.RoleName)
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
