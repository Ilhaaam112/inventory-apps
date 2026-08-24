package repository

import "database/sql"

// LaporanRepository menampung seluruh query laporan.
// Semua laporan hanya membaca data, tidak pernah mengubah stok,
// jadi cukup satu repository dengan method yang dipisah per file.
type LaporanRepository struct {
	db *sql.DB
}

func NewLaporanRepository(db *sql.DB) *LaporanRepository {
	return &LaporanRepository{db: db}
}
