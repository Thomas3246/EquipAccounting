package sqlite

import (
	"context"
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	_ "modernc.org/sqlite"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users(login, password, name, isAdmin, department) VALUES (?,?,?,?,?)"
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Password, user.Name, user.IsAdmin, user.DepartmentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	query := "SELECT id, password, name, isAdmin, department FROM users WHERE login = ?"

	row := r.db.QueryRowContext(ctx, query, login)

	user := &domain.User{Login: login}

	err := row.Scan(&user.Id, &user.Password, &user.Name, &user.IsAdmin, &user.DepartmentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
