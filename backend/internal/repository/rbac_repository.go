package repository

import "database/sql"

// RBACRepository memisahkan pertanyaan "kamu boleh apa" dari
// pertanyaan "kamu siapa" yang ditangani UserRepository.
type RBACRepository struct {
	db *sql.DB
}

func NewRBACRepository(db *sql.DB) *RBACRepository {
	return &RBACRepository{db: db}
}

// Permissions mengembalikan seluruh kode permission milik satu user,
// diturunkan dari role-nya.
func (r *RBACRepository) Permissions(userID int) []string {
	rows, err := r.db.Query(`
		SELECT p.code
		FROM users u
		JOIN role_permissions rp ON rp.role_id = u.role_id
		JOIN permissions p       ON p.id       = rp.permission_id
		WHERE u.id = ?
		ORDER BY p.code`, userID)
	if err != nil {
		return []string{}
	}
	defer rows.Close()

	list := []string{}
	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err == nil {
			list = append(list, code)
		}
	}
	return list
}
