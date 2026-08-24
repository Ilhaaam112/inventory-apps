package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

// rowScanner menampung *sql.Row maupun *sql.Rows, supaya satu fungsi scan
// bisa dipakai untuk GetAll dan GetByID sekaligus.
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// ensureExists memastikan sebuah baris master data benar-benar ada.
func ensureExists(tx *sql.Tx, table string, id int, pesan string) error {
	var x int
	err := tx.QueryRow(fmt.Sprintf("SELECT id FROM %s WHERE id = ?", table), id).Scan(&x)
	if err == sql.ErrNoRows {
		return errors.New(pesan)
	}
	return err
}

// generateCode membuat nomor transaksi berurutan, misal IN-0001, TRF-0007.
func generateCode(tx *sql.Tx, table, prefix string) (string, error) {
	var last sql.NullString
	q := fmt.Sprintf("SELECT MAX(code) FROM %s WHERE code LIKE ?", table)
	if err := tx.QueryRow(q, prefix+"-%").Scan(&last); err != nil {
		return "", err
	}
	num := 1
	if last.Valid && len(last.String) > len(prefix)+1 {
		var n int
		if _, err := fmt.Sscanf(last.String[len(prefix)+1:], "%d", &n); err == nil {
			num = n + 1
		}
	}
	return fmt.Sprintf("%s-%04d", prefix, num), nil
}

// currentStock membaca stok sekarang sekaligus mengunci barisnya sampai commit.
func currentStock(tx *sql.Tx, lokasiID, barangID int) (int, error) {
	var qty int
	err := tx.QueryRow(
		"SELECT quantity FROM warehouse_stocks WHERE lokasi_id = ? AND barang_id = ? FOR UPDATE",
		lokasiID, barangID,
	).Scan(&qty)
	if err == sql.ErrNoRows {
		if _, err := tx.Exec(
			"INSERT INTO warehouse_stocks (lokasi_id, barang_id, quantity) VALUES (?, ?, 0)",
			lokasiID, barangID,
		); err != nil {
			return 0, err
		}
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return qty, nil
}

// applyMovement adalah SATU-SATUNYA pintu untuk mengubah stok.
// delta positif menambah stok, delta negatif mengurangi.
// Fungsi ini sekaligus menolak transaksi yang membuat stok minus.
func applyMovement(tx *sql.Tx, lokasiID, barangID, delta int, moveType, refType string, refID int) error {
	before, err := currentStock(tx, lokasiID, barangID)
	if err != nil {
		return err
	}

	after := before + delta
	if after < 0 {
		var nama string
		tx.QueryRow("SELECT nama FROM barang WHERE id = ?", barangID).Scan(&nama)
		return fmt.Errorf("stok %s di gudang tidak mencukupi (tersedia %d, diminta %d)", nama, before, -delta)
	}

	if _, err := tx.Exec(
		"UPDATE warehouse_stocks SET quantity = ? WHERE lokasi_id = ? AND barang_id = ?",
		after, lokasiID, barangID,
	); err != nil {
		return err
	}

	if _, err := tx.Exec(`
		INSERT INTO stock_movements
			(barang_id, lokasi_id, type, reference_type, reference_id, quantity, stock_before, stock_after)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		barangID, lokasiID, moveType, refType, refID, delta, before, after,
	); err != nil {
		return err
	}

	// Sinkronkan kolom barang.stok sebagai total seluruh gudang,
	// supaya halaman Data Barang tetap menampilkan angka yang benar.
	_, err = tx.Exec(`
		UPDATE barang
		SET stok = (SELECT COALESCE(SUM(quantity), 0) FROM warehouse_stocks WHERE barang_id = ?)
		WHERE id = ?`, barangID, barangID)
	return err
}
