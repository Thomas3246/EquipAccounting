package handler

import (
	"fmt"
	"net/http"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: *service}
}

func (h *UserHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logged")
}
