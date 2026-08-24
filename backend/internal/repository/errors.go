package repository

import (
	"errors"
	"strings"
)

// Error yang aman ditampilkan ke user. Pesan asli dari MySQL tidak
// pernah diteruskan, karena bisa memuat nama tabel, nama kolom, dan
// potongan query yang berguna bagi penyerang.
var (
	ErrDuplikat = errors.New("data dengan nama tersebut sudah ada")
	ErrDipakai  = errors.New("data tidak bisa dihapus karena masih dipakai data lain")
	ErrSimpan   = errors.New("gagal menyimpan data")
	ErrHapus    = errors.New("gagal menghapus data")
)

// bungkusError menerjemahkan error driver MySQL menjadi pesan umum.
func bungkusError(err error) error {
	if err == nil {
		return nil
	}
	pesan := err.Error()
	switch {
	case strings.Contains(pesan, "Duplicate entry"):
		return ErrDuplikat
	case strings.Contains(pesan, "foreign key constraint"),
		strings.Contains(pesan, "Cannot delete"),
		strings.Contains(pesan, "a foreign key"):
		return ErrDipakai
	}
	return ErrSimpan
}
