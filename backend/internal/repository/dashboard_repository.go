package repository

import "database/sql"

// DashboardRepository menampung query ringkasan untuk halaman utama.
// Semuanya hanya membaca, tidak pernah mengubah stok.
type DashboardRepository struct {
	db *sql.DB
}

func NewDashboardRepository(db *sql.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}
