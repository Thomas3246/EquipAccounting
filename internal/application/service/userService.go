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

func (s *UserService) Authenticate(login string, password string) (*domain.User, error) {
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

func (s *UserService) Register(login string, password string, role string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if login == "" || password == "" || role == "" {
		return ErrNullParameter
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user := &domain.User{
		Login:    login,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
