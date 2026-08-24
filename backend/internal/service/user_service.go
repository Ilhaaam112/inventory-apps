package service

import (
	"errors"

	"github.com/username/belajar_go/backend/internal/model"
	"github.com/username/belajar_go/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Login(username, password string) (model.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return model.User{}, errors.New("username atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, errors.New("username atau password salah")
	}

	return user, nil
}



// ---------- BARU: untuk halaman Profil ----------

func (s *UserService) GetProfile(id int) (model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateProfile(id int, namaLengkap string) (model.User, error) {
	if namaLengkap == "" {
		return model.User{}, errors.New("nama lengkap tidak boleh kosong")
	}
	if err := s.repo.UpdateProfile(id, namaLengkap); err != nil {
		return model.User{}, err
	}
	return s.repo.GetByID(id)
}

func (s *UserService) ChangePassword(id int, oldPassword, newPassword string) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("password lama salah")
	}
	if len(newPassword) < 6 {
		return errors.New("password baru minimal 6 karakter")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.repo.UpdatePassword(id, string(hashed))
}