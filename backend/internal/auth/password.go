package auth

import "golang.org/x/crypto/bcrypt"

// Cost 12 (bawaan bcrypt adalah 10). Lebih lambat sedikit saat login,
// tetapi jauh lebih mahal untuk ditebak secara brute force.
const bcryptCost = 12

// dummyHash dipakai saat username tidak ditemukan, supaya waktu respons
// login mirip dengan kasus password salah. Tanpa ini, penyerang bisa
// menebak username mana yang ada hanya dari selisih waktu balasan.
const dummyHash = "$2a$12$C6UzMDM.H6dfI/f/IKcEe.NNvVcxJH/hcxRW.p9dCRc7iCJlEfN0O"

func HashPassword(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcryptCost)
	return string(b), err
}

func CheckPassword(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}

// BurnTime sengaja membakar waktu setara satu pengecekan bcrypt.
func BurnTime(plain string) {
	_ = bcrypt.CompareHashAndPassword([]byte(dummyHash), []byte(plain))
}
