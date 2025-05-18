package sqlite

import (
	"database/sql"
	"fmt"

	domain "github.com/Thomas3246/EquipAccounting/internal/domain/user"
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
