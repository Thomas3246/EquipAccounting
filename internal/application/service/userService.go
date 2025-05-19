package service

import (
	"context"
	"time"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/user"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo database.UserRepo
}

func NewUserService(repo database.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Authenticate(login, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	user, err := s.repo.GetByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
