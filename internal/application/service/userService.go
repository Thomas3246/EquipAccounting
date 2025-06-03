package service

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
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

func (s *UserService) Register(login string, password string, name string, isAdmin string, department string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if login == "" || password == "" || name == "" || isAdmin == "" || department == "" {
		return ErrNullParameter
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	intIsAdmin, err := strconv.Atoi(isAdmin)
	if err != nil {
		return ErrInvalidParameter
	}

	intDepartment, err := strconv.Atoi(department)
	if err != nil {
		return ErrInvalidParameter
	}

	user := &domain.User{
		Login:        login,
		Password:     string(hashedPassword),
		Name:         name,
		IsAdmin:      intIsAdmin,
		DepartmentId: intDepartment,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserByLogin(login string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	user, err := s.repo.GetByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsers() ([]domain.ViewUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := s.repo.GetUsersView(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if user.IsAdmin == 1 {
		return ErrNoAccess
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) AddUser(user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	err = s.repo.New(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserById(id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (s *UserService) EditUser(user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.UpdateUserData(ctx, user)
	if err != nil {
		return err
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return err
		}
		err = s.repo.ChangeUserPassword(ctx, user.Id, string(hashedPassword))
		if err != nil {
			return err
		}
	}

	return nil

}
