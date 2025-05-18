package service

import "github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"

type UserService struct {
	repo database.UserRepo
}

func NewUserService(repo database.UserRepo) *UserService {
	return &UserService{repo: repo}
}
