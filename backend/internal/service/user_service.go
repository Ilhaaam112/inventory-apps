package service

import (
	"errors"
	"strings"

	"github.com/username/belajar_go/backend/internal/auth"
	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(id int) (model.User, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return model.User{}, err
	}
	u.Password = ""
	return u, nil
}

func (s *UserService) UpdateProfile(id int, namaLengkap string) (model.User, error) {
	namaLengkap = strings.TrimSpace(namaLengkap)
	if namaLengkap == "" {
		return model.User{}, errors.New("nama lengkap tidak boleh kosong")
	}
	if len(namaLengkap) > 100 {
		return model.User{}, errors.New("nama lengkap maksimal 100 karakter")
	}
	if err := s.repo.UpdateProfile(id, namaLengkap); err != nil {
		return model.User{}, errors.New("gagal menyimpan profil")
	}
	return s.GetProfile(id)
}

func (s *UserService) ChangePassword(id int, oldPassword, newPassword string) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	if !auth.CheckPassword(user.Password, oldPassword) {
		return errors.New("password lama salah")
	}
	if len(newPassword) < 8 {
		return errors.New("password baru minimal 8 karakter")
	}
	if len(newPassword) > 200 {
		return errors.New("password baru terlalu panjang")
	}
	if newPassword == oldPassword {
		return errors.New("password baru harus berbeda dari password lama")
	}

	hashed, err := auth.HashPassword(newPassword)
	if err != nil {
		return errors.New("gagal mengubah password")
	}
	return s.repo.UpdatePassword(id, hashed)
}
