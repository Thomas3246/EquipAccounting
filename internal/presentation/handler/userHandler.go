package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Thomas3246/EquipAccounting/internal/application/service"
	templateloader "github.com/Thomas3246/EquipAccounting/pkg/templateLoader"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: *service}
}

func (h *UserHandler) LoginGet(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(templateloader.GetTemplatePath("loginPage.html"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error by parse", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error by execute")
	}
}

func (h *UserHandler) LoginPost(w http.ResponseWriter, r *http.Request) {

}
