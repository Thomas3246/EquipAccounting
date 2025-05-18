package database

import domain "github.com/Thomas3246/EquipAccounting/internal/domain/user"

type UserRepo interface {
	Create(*domain.User) error
}
