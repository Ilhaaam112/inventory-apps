package service

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Validasi input dilakukan di sini, tidak bergantung pada React.
// Siapa pun bisa mengirim POST langsung ke API tanpa lewat browser.

func teksWajib(nilai, label string, maks int) (string, error) {
	nilai = strings.TrimSpace(nilai)
	if nilai == "" {
		return "", fmt.Errorf("%s wajib diisi", label)
	}
	if utf8.RuneCountInString(nilai) > maks {
		return "", fmt.Errorf("%s maksimal %d karakter", label, maks)
	}
	return nilai, nil
}

func teksOpsional(nilai, label string, maks int) (string, error) {
	nilai = strings.TrimSpace(nilai)
	if utf8.RuneCountInString(nilai) > maks {
		return "", fmt.Errorf("%s maksimal %d karakter", label, maks)
	}
	return nilai, nil
}

func idValid(id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	return nil
}
