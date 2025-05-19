package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/user"
	_ "modernc.org/sqlite"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *domain.User) error {
	fmt.Println("Added")
	return nil
}

func (r *UserRepo) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	query := "SELECT id, password, role FROM users WHERE login = ?"

	row := r.db.QueryRowContext(ctx, query, login)

	user := &domain.User{Login: login}

	err := row.Scan(&user.Id, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
